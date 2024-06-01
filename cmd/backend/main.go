package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/a-h/templ"
	"github.com/gorilla/websocket"
	"github.com/kevinkjt2000/factory-planner/components"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}
	// Intentionally keep connection open forever doing nothing
	// when the server exits, live reload will trigger in the web browser
	// defer c.Close()
}

func main() {
	quitOnSigTerm()
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	fmt.Println("Registering handlers")
	http.Handle("/", templ.Handler(components.Page()))
	http.Handle("/ws", webSocketHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to listen and serve: %v\n", err)
	}
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
