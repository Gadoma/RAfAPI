package http

import (
	applicationAff "github.com/gadoma/rafapi/internal/affirmation/application"

	dbAff "github.com/gadoma/rafapi/internal/affirmation/infrastructure/database"

	dbCommon "github.com/gadoma/rafapi/internal/common/infrastructure/database"
	httpCommon "github.com/gadoma/rafapi/internal/common/infrastructure/http"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) BootstrapControllers(db *dbCommon.DB) []httpCommon.Controller {
	responder := httpCommon.NewResponder()

	rafService := applicationAff.NewAffirmationService(dbAff.NewAffirmationRepository(db))

	rafReqHandler := NewAffirmationRequestHandler()

	return []httpCommon.Controller{
		NewAffirmationController(rafService, responder, rafReqHandler),
	}
}
