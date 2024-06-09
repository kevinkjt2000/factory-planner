package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
	"github.com/kevinkjt2000/factory-planner/components"
	"github.com/kevinkjt2000/factory-planner/internal"
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

func pageHandler(w http.ResponseWriter, r *http.Request) {
	db := internal.NewDatabase()
	err := components.Page(db.GetRecipeTypes()).Render(r.Context(), w)
	if err != nil {
		fmt.Printf("error during page render: %v", err)
	}
}

func main() {
	quitOnSigTerm()
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	fmt.Println("Registering handlers")
	http.HandleFunc("/", pageHandler)
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
