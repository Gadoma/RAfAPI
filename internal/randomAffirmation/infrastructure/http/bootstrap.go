package http

import (
	applicationRaf "github.com/gadoma/rafapi/internal/randomAffirmation/application"

	dbRaf "github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/database"

	dbCommon "github.com/gadoma/rafapi/internal/common/infrastructure/database"
	httpCommon "github.com/gadoma/rafapi/internal/common/infrastructure/http"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) BootstrapControllers(db *dbCommon.DB) []httpCommon.Controller {
	responder := httpCommon.NewResponder()

	rafService := applicationRaf.NewRandomAffirmationService(dbRaf.NewRandomAffirmationRepository(db))

	rafReqHandler := NewRandomAffirmationRequestHandler()

	return []httpCommon.Controller{
		NewRandomAffirmationController(rafService, responder, rafReqHandler),
	}
}
