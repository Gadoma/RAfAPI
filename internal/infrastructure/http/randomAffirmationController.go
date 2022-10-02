package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerRandomAffirmationRoutes(r *mux.Router) {
	r.HandleFunc("/random_affirmation", s.handleGetRandomAffirmation).Methods("GET").Name("getRandomAffirmation")
}

func (s *Server) handleGetRandomAffirmation(w http.ResponseWriter, r *http.Request) {
	categoryIds, err := s.getRandomAffirmationCategoryIdsParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	affirmation, err := s.RandomAffirmationService.GetRandomAffirmation(r.Context(), categoryIds)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count := 0

	if affirmation.Text != "" {
		count = 1
	}

	s.respondSuccessOk(w, affirmation, count)
}
