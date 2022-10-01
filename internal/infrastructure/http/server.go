package http

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gorilla/mux"
)

const ShutdownTimeout = 1 * time.Second

type Server struct {
	ln     net.Listener
	server *http.Server
	router *mux.Router

	Addr   string
	Domain string

	AffirmationService domain.AffirmationService
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	s.router.Use(s.handlePanicMiddleware)

	s.server.Handler = http.HandlerFunc(s.serveHTTP)

	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound)

	r := s.router.PathPrefix("/").Subrouter()

	s.registerAffirmationRoutes(r)

	return s
}

func (s *Server) Open() (err error) {
	if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
		return err
	}

	go s.server.Serve(s.ln)

	return nil
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleNotFound(w http.ResponseWriter, r *http.Request) {
	s.respondError(w, "Not found", http.StatusNotFound)
}

func (s *Server) parseIdParameter(p string) (uint64, error) {
	id, err := strconv.ParseUint(p, 10, 64)

	if err != nil || id == 0 {
		return 0, err
	}

	return id, nil
}
