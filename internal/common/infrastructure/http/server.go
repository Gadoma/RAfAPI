package http

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const ShutdownTimeout = 1 * time.Second

type Server struct {
	ln        net.Listener
	server    *http.Server
	router    *mux.Router
	responder *Responder

	Addr   string
	Domain string

	controllers []Controller
}

func NewServer(controllers []Controller) *Server {
	s := &Server{
		server: &http.Server{
			ReadTimeout:       1 * time.Second,
			WriteTimeout:      1 * time.Second,
			IdleTimeout:       30 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
		},
		router:      mux.NewRouter(),
		responder:   NewResponder(),
		controllers: controllers,
	}

	s.router.Use(s.handlePanicMiddleware)

	s.server.Handler = http.HandlerFunc(s.serveHTTP)

	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound)

	s.RegisterRoutes()

	return s
}

func (s *Server) RegisterRoutes() {
	r := s.router.PathPrefix("/").Subrouter()

	for _, c := range s.controllers {
		c.RegisterRoutes(r)
	}
}

func (s *Server) Open() (err error) {
	if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
		return err
	}

	serve := func() {
		err = s.server.Serve(s.ln)
	}

	go serve()

	return
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleNotFound(w http.ResponseWriter, _ *http.Request) {
	s.responder.RespondError(w, "Not found", http.StatusNotFound)
}
