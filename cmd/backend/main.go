package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevinkjt2000/factory-planner/internal"
)

/* To-Do List
Anywhere that has a long load time needs an svg spinner from http://samherbert.net/svg-loaders/
Fully adopt entgo as a replacement for gorm (which shockingly lacks the ability to do multiple joins!!!!)
*/

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
