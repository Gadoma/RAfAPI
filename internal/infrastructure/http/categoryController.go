package http

import (
	"net/http"

	"github.com/gadoma/rafapi/internal/domain"
	"github.com/gorilla/mux"
)

func (s *Server) registerCategoryRoutes(r *mux.Router) {
	r.HandleFunc("/categories", s.handleGetCategories).Methods("GET").Name("getCategories")
	r.HandleFunc("/categories", s.handleCreateCategory).Methods("POST").Name("createCategory")
	r.HandleFunc("/categories/{categoryId:[0-9]+}", s.handleGetCategory).Methods("GET").Name("getCategory")
	r.HandleFunc("/categories/{categoryId:[0-9]+}", s.handleUpdateCategory).Methods("PUT").Name("updateCategory")
	r.HandleFunc("/categories/{categoryId:[0-9]+}", s.handleDeleteCategory).Methods("DELETE").Name("deleteCategory")
}

func (s *Server) handleGetCategories(w http.ResponseWriter, r *http.Request) {
	categories, n, err := s.CategoryService.GetCategories(r.Context())

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessOk(w, categories, n)
}

func (s *Server) handleGetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := s.getCategoryIdParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	categories, err := s.CategoryService.GetCategory(r.Context(), id)

	if err == domain.ErrorResourceNotFound {
		s.respondErrorNotFound(w)
		return
	}

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessOk(w, categories, 1)
}

func (s *Server) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	au, err := s.getCategoryUpdate(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.CategoryService.CreateCategory(r.Context(), au)

	if err == domain.ErrorCategoryUpdateInvalidName {
		s.respondError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessOk(w, id, 1)
}

func (s *Server) handleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := s.getCategoryIdParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	au, err := s.getCategoryUpdate(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.CategoryService.UpdateCategory(r.Context(), id, au)

	if err == domain.ErrorCategoryUpdateInvalidName {
		s.respondError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessNoContent(w)
}

func (s *Server) handleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := s.getCategoryIdParameter(r)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.CategoryService.DeleteCategory(r.Context(), id)

	if err != nil {
		s.respondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.respondSuccessNoContent(w)
}
