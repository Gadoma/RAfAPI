package http

import (
	"errors"
	"net/http"

	"github.com/gadoma/rafapi/internal/affirmation/domain"
	common "github.com/gadoma/rafapi/internal/common/domain"
	"github.com/gorilla/mux"
)

type AffirmationController struct {
	service    domain.AffirmationService
	responder  AffirmationResponder
	reqHandler *AffirmationRequestHandler
}

func NewAffirmationController(service domain.AffirmationService, responder AffirmationResponder, reqHandler *AffirmationRequestHandler) *AffirmationController {
	return &AffirmationController{
		service:    service,
		responder:  responder,
		reqHandler: reqHandler,
	}
}

func (c *AffirmationController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/affirmations", c.handleGetAffirmations).Methods("GET").Name("getAffirmations")
	r.HandleFunc("/affirmations", c.handleCreateAffirmation).Methods("POST").Name("createAffirmation")
	r.HandleFunc("/affirmations/{affirmationId:[0-7][0-9A-HJKMNP-TV-Z]{25}}", c.handleGetAffirmation).Methods("GET").Name("getAffirmation")
	r.HandleFunc("/affirmations/{affirmationId:[0-7][0-9A-HJKMNP-TV-Z]{25}}", c.handleUpdateAffirmation).Methods("PUT").Name("updateAffirmation")
	r.HandleFunc("/affirmations/{affirmationId:[0-7][0-9A-HJKMNP-TV-Z]{25}}", c.handleDeleteAffirmation).Methods("DELETE").Name("deleteAffirmation")
}

func (c *AffirmationController) handleGetAffirmations(w http.ResponseWriter, r *http.Request) {
	affirmations, n, err := c.service.GetAffirmations(r.Context())

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.responder.RespondSuccessOk(w, affirmations, n)
}

func (c *AffirmationController) handleGetAffirmation(w http.ResponseWriter, r *http.Request) {
	id, err := c.reqHandler.getAffirmationIdParameter(r)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	affirmations, err := c.service.GetAffirmation(r.Context(), *id)

	if errors.Is(err, common.ErrorResourceNotFound) {
		c.responder.RespondErrorNotFound(w)
		return
	}

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.responder.RespondSuccessOk(w, affirmations, 1)
}

func (c *AffirmationController) handleCreateAffirmation(w http.ResponseWriter, r *http.Request) {
	cac, err := c.reqHandler.getCreateAffirmationCommand(r)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.service.CreateAffirmation(r.Context(), cac)

	if errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidId) || errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidCategoryId) || errors.Is(err, domain.ErrorCreateAffirmationCommandInvalidText) {
		c.responder.RespondError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.responder.RespondSuccessOk(w, id, 1)
}

func (c *AffirmationController) handleUpdateAffirmation(w http.ResponseWriter, r *http.Request) {
	id, err := c.reqHandler.getAffirmationIdParameter(r)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	uac, err := c.reqHandler.getUpdateAffirmationCommand(r)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.UpdateAffirmation(r.Context(), *id, uac)

	if errors.Is(err, domain.ErrorUpdateAffirmationCommandInvalidCategoryId) || errors.Is(err, domain.ErrorUpdateAffirmationCommandInvalidText) {
		c.responder.RespondError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.responder.RespondSuccessNoContent(w)
}

func (c *AffirmationController) handleDeleteAffirmation(w http.ResponseWriter, r *http.Request) {
	id, err := c.reqHandler.getAffirmationIdParameter(r)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.DeleteAffirmation(r.Context(), *id)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.responder.RespondSuccessNoContent(w)
}
