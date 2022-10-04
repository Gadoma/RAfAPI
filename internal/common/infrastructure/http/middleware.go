package http

import (
	"fmt"
	"net/http"
)

func (s *Server) handlePanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.responder.RespondError(w, fmt.Sprintf("a general error occurred: %q", err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
