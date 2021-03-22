package main

import (
	"context"
	"flag"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultDaprPort    = 3500
	defaultServicePort = 50001
	defaultStoreName   = "statestore"
)

var (
	storeName   string
	servicePort int
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("client new begin")
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	fmt.Println("client new success")
	go func() {
		for {
			fmt.Println("pong")
			select {
			case <-time.After(time.Second):
				if resp, err := client.InvokeMethodWithContent(ctx, "app-ping", "ping", "get", &dapr.DataContent{}); err != nil {
					panic(err)
				} else {
					fmt.Printf("response: %+v", string(resp))
				}
			}
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGSTOP)
	<-sign
}

func init() {
	flag.StringVar(&storeName, "storeName", defaultStoreName, "storeName name")
	flag.IntVar(&servicePort, "service-port", defaultServicePort, "service port")
	flag.Parse()
}
