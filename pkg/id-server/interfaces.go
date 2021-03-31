package id_server

import (
	"context"
	"github.com/imfuxiao/dapr-tutorial/pkg/id-server/v1"
)

type MachineIdGenerator interface {
	MachineId(ctx context.Context) (int64, error)
	MachineIdByIP(ctx context.Context, ip string) (int64, error)
}

type IdGenerator interface {
	NewId() int64
	ExplainId(id int64) v1.ID
	ManualId(time, seq, machineId int64) int64
}
