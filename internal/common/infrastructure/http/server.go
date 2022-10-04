package http

import (
	"context"
	"net"
	"net/http"
	"time"

	httpAff "github.com/gadoma/rafapi/internal/affirmation/infrastructure/http"
	httpCat "github.com/gadoma/rafapi/internal/category/infrastructure/http"
	httpRaf "github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/http"

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

	AffirmationController       *httpAff.AffirmationController
	CategoryController          *httpCat.CategoryController
	RandomAffirmationController *httpRaf.RandomAffirmationController
}

func NewServer() *Server {
	s := &Server{
		server:    &http.Server{},
		router:    mux.NewRouter(),
		responder: NewResponder(),
	}

	s.router.Use(s.handlePanicMiddleware)

	s.server.Handler = http.HandlerFunc(s.serveHTTP)

	s.router.NotFoundHandler = http.HandlerFunc(s.handleNotFound)

	return s
}

func (s *Server) RegisterRoutes() {
	r := s.router.PathPrefix("/").Subrouter()

	s.AffirmationController.RegisterAffirmationRoutes(r)
	s.CategoryController.RegisterCategoryRoutes(r)
	s.RandomAffirmationController.RegisterRandomAffirmationRoutes(r)
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
	s.responder.RespondError(w, "Not found", http.StatusNotFound)
}
