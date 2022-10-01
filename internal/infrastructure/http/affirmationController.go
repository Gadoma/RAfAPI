package http

import (
	"net/http"

	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gorilla/mux"
)

func (s *Server) registerAffirmationRoutes(r *mux.Router) {
	r.HandleFunc("/affirmations", s.handleGetAffirmations).Methods("GET").Name("getAffirmations")
	r.HandleFunc("/affirmations", s.handleCreateAffirmation).Methods("POST").Name("createAffirmation")
	r.HandleFunc("/affirmations/{affirmationId:[0-9]+}", s.handleGetAffirmation).Methods("GET").Name("getAffirmation")
	r.HandleFunc("/affirmations/{affirmationId:[0-9]+}", s.handleUpdateAffirmation).Methods("PUT").Name("updateAffirmation")
	r.HandleFunc("/affirmations/{affirmationId:[0-9]+}", s.handleDeleteAffirmation).Methods("DELETE").Name("deleteAffirmation")
}

func (s *Server) handleGetAffirmations(w http.ResponseWriter, r *http.Request) {
	affirmations, n, err := s.AffirmationService.GetAffirmations(r.Context())

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessOk(w, affirmations, n)
}

func (s *Server) handleGetAffirmation(w http.ResponseWriter, r *http.Request) {
	id, err := s.getAffirmationIdParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	affirmations, err := s.AffirmationService.GetAffirmation(r.Context(), id)

	if err == domain.ErrorResourceNotFound {
		s.respondErrorNotFound(w)
		return
	}

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessOk(w, affirmations, 1)
}

func (s *Server) handleCreateAffirmation(w http.ResponseWriter, r *http.Request) {
	au, err := s.getAffirmationUpdate(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.AffirmationService.CreateAffirmation(r.Context(), au)

	if err == domain.ErrorAffirmationUpdateInvalidCategoryId || err == domain.ErrorAffirmationUpdateInvalidText {
		s.respondError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessOk(w, id, 1)
}

func (s *Server) handleUpdateAffirmation(w http.ResponseWriter, r *http.Request) {
	id, err := s.getAffirmationIdParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	au, err := s.getAffirmationUpdate(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AffirmationService.UpdateAffirmation(r.Context(), id, au)

	if err == domain.ErrorAffirmationUpdateInvalidCategoryId || err == domain.ErrorAffirmationUpdateInvalidText {
		s.respondError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessNoContent(w)
}

func (s *Server) handleDeleteAffirmation(w http.ResponseWriter, r *http.Request) {
	id, err := s.getAffirmationIdParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AffirmationService.DeleteAffirmation(r.Context(), id)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessNoContent(w)
}
