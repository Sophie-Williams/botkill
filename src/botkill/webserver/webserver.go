package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Server struct holds information about the server, like
// the interface it's binded, port it's using and the games it's
// hosting.
type Server struct {
	//Games  []*game.Game `json:"games"`
	Host   string      `json:"host"`
	Port   int         `json:"port"`
	Router *mux.Router `json:"router"`
}

// NewServer will start the http server, that will listen to port 8080,
// and return the server or error
func NewServer(host string, port string) *Server {
	s := &Server{}

	// Init a new Gorilla Mux router
	s.Router = mux.NewRouter()

	// Set the root "/" url to serve static files from www directory
	s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))

	http.Handle("/", s.Router)
	go http.ListenAndServe(":8080", nil)

	return s
}
