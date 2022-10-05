package http

import (
	"net/http"

	"github.com/gadoma/rafapi/internal/randomAffirmation/domain"
	"github.com/gorilla/mux"
)

type RandomAffirmationController struct {
	service    domain.RandomAffirmationService
	responder  RandomAffirmationResponder
	reqHandler *RandomAffirmationRequestHandler
}

func NewRandomAffirmationController(service domain.RandomAffirmationService, responder RandomAffirmationResponder, reqHandler *RandomAffirmationRequestHandler) *RandomAffirmationController {
	return &RandomAffirmationController{
		service:    service,
		responder:  responder,
		reqHandler: reqHandler,
	}
}

func (c *RandomAffirmationController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/random_affirmation", c.handleGetRandomAffirmation).Methods("GET").Name("getRandomAffirmation")
}

func (c *RandomAffirmationController) handleGetRandomAffirmation(w http.ResponseWriter, r *http.Request) {
	categoryIds, err := c.reqHandler.getRandomAffirmationCategoryIdsParameter(r)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusBadRequest)
		return
	}

	affirmation, err := c.service.GetRandomAffirmation(r.Context(), categoryIds)

	if err != nil {
		c.responder.RespondError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count := 0

	if affirmation.Text != "" {
		count = 1
	}

	c.responder.RespondSuccessOk(w, affirmation, count)
}
