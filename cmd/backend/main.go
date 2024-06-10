package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinkjt2000/factory-planner/internal"
)

func main() {
	quitOnSigTerm()
	internal.StartHttpServer()
}

func quitOnSigTerm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(0)
	}()
}
