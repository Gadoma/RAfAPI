package http

import (
	applicationAff "github.com/gadoma/rafapi/internal/affirmation/application"
	applicationCat "github.com/gadoma/rafapi/internal/category/application"
	appicationRaf "github.com/gadoma/rafapi/internal/randomAffirmation/application"

	httpAff "github.com/gadoma/rafapi/internal/affirmation/infrastructure/http"
	httpCat "github.com/gadoma/rafapi/internal/category/infrastructure/http"
	httpRaf "github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/http"

	dbAff "github.com/gadoma/rafapi/internal/affirmation/infrastructure/database"
	dbCat "github.com/gadoma/rafapi/internal/category/infrastructure/database"
	dbRaf "github.com/gadoma/rafapi/internal/randomAffirmation/infrastructure/database"

	dbCommon "github.com/gadoma/rafapi/internal/common/infrastructure/database"
)

type Bootstrapper interface {
	BootstrapControllers(db *dbCommon.DB) []Controller
}

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) BootstrapControllers(db *dbCommon.DB) []Controller {
	responder := NewResponder()

	affService := applicationAff.NewAffirmationService(dbAff.NewAffirmationRepository(db))
	catService := applicationCat.NewCategoryService(dbCat.NewCategoryRepository(db))
	rafService := appicationRaf.NewRandomAffirmationService(dbRaf.NewRandomAffirmationRepository(db))

	affReqHandler := httpAff.NewAffirmationRequestHandler()
	catReqHandler := httpCat.NewCategoryRequestHandler()
	rafReqHandler := httpRaf.NewRandomAffirmationRequestHandler()

	return []Controller{
		httpAff.NewAffirmationController(affService, responder, affReqHandler),
		httpCat.NewCategoryController(catService, responder, catReqHandler),
		httpRaf.NewRandomAffirmationController(rafService, responder, rafReqHandler),
	}
}
