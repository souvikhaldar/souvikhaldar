package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	router *mux.Router
}

func NewServer(r *mux.Router) *Server {
	return &Server{
		router: r,
	}
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("-------------------------------------")
		log.Info("Serving the host: ", r.RemoteAddr)
		log.Info("Serving the agent: ", r.Header.Get("User-Agent"))
		http.ServeFile(w, r, "./static/index.html")
	}
}

func (s *Server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Ping request")
		fmt.Fprint(w, "pong")
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func main() {
	// create the dependencies required from the server

	// create a new mux router
	r := mux.NewRouter()

	// create a new server
	s := NewServer(r)

	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/ping", s.handlePing())

	http.ListenAndServe(":80", s)
}
