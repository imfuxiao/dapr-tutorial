package id_server

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/imfuxiao/dapr-tutorial/pkg/util"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

type MachineIdMaxError struct {
}

func (m MachineIdMaxError) Error() string {
	return fmt.Sprintf("Exceeding the maximum machine_id = %d", MaxMachineSeq)
}

type MachineIdNotFoundPodIpError struct {
}

func (m MachineIdNotFoundPodIpError) Error() string {
	return "Pod ip not found"
}

type MachineDaprState struct {
	client dapr.Client
}

func (m *MachineDaprState) MachineId(ctx context.Context) (int64, error) {
	ip, exists := getPodIp()
	if !exists {
		return 0, MachineIdNotFoundPodIpError{}
	}
	return m.MachineIdByIP(ctx, ip)
}

func (m *MachineDaprState) MachineIdByIP(ctx context.Context, ip string) (int64, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	machineIdKey := fmt.Sprintf(MachineIdKeyFmt, ip)
	item, err := m.client.GetState(ctx, storeName, machineIdKey)
	if err == nil {
		value, err := util.NewStateItem(item).GetIntValue()
		if err != nil {
			return 0, errors.Wrapf(err, "%s = %s, to int error", machineIdKey, string(item.Value))
		}
		return int64(value), nil
	}

	fmt.Printf("get %s error: %s\n", machineIdKey, err.Error())

	// CAS
	item, err = m.client.GetState(ctx, storeName, MachineIdSeq)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item)

	seq, _ := strconv.Atoi(string(item.Value))
	if seq >= MaxMachineSeq {
		return 0, MachineIdMaxError{}
	}

	option := &dapr.StateOptions{
		Concurrency: dapr.StateConcurrencyFirstWrite,
		Consistency: dapr.StateConsistencyStrong,
	}

	machineIdSeqOpt := &dapr.StateOperation{
		Type: dapr.StateOperationTypeUpsert,
		Item: &dapr.SetStateItem{
			Key:     item.Key,
			Value:   nil,
			Etag:    &dapr.ETag{Value: ""},
			Options: option,
		},
	}

	_ = m.client.ExecuteStateTransaction(ctx, storeName, map[string]string{}, []*dapr.StateOperation{machineIdSeqOpt})
	return 0, nil
}

func getPodIp() (string, bool) {
	return os.LookupEnv(EnvPodIp)
}
