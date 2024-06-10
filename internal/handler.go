package internal

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/kevinkjt2000/factory-planner/components"
)

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type App struct {
	IndexHandler     *IndexHandler
	SearchHandler    *SearchHandler
	WebsocketHandler *WebsocketHandler
}

func (h *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request received: %s %s\n", req.Method, req.URL.String())
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	switch head {
	case "":
		h.IndexHandler.ServeHTTP(w, req)
	case "search":
		h.SearchHandler.ServeHTTP(w, req)
	case "ws":
		h.WebsocketHandler.ServeHTTP(w, req)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

type SearchHandler struct{}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.handleGet(w, req)
	default:
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
	}
}

func (h *SearchHandler) handleGet(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	db := NewDatabase()
	defer db.Close()

	recipes := db.SearchRecipeItemOutputs(item)
	err := components.SearchResults(recipes).Render(req.Context(), w)
	if err != nil {
		fmt.Printf("searchHandler render: %v", err)
	}
}

type WebsocketHandler struct {
	upgrader websocket.Upgrader
}

func (h *WebsocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, err := h.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
	}
	// Intentionally keep connection open forever doing nothing
	// when the server exits, live reload will trigger in the web browser
	// defer c.Close()
}

type IndexHandler struct{}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	db := NewDatabase()
	defer db.Close()
	content := components.Page(db.GetRecipeTypes())
	err := components.Layout(content).Render(req.Context(), w)
	if err != nil {
		fmt.Printf("error during page render: %v", err)
	}
}

func StartHttpServer() {
	fmt.Println("Registering handlers")
	app := &App{
		SearchHandler:    new(SearchHandler),
		WebsocketHandler: new(WebsocketHandler),
	}

	fmt.Println("Starting http service")
	err := http.ListenAndServe(":8080", app)
	if err != nil {
		fmt.Printf("failed to listen and serve: %v\n", err)
	}
}
