package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	go func() {
		if err := r.Run(fmt.Sprintf(":%d", servicePort)); err != nil {
			fmt.Printf("run error: %#v\n", err)
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
