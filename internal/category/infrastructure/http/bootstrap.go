package http

import (
	applicationCat "github.com/gadoma/rafapi/internal/category/application"

	dbCat "github.com/gadoma/rafapi/internal/category/infrastructure/database"

	dbCommon "github.com/gadoma/rafapi/internal/common/infrastructure/database"
	httpCommon "github.com/gadoma/rafapi/internal/common/infrastructure/http"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) BootstrapControllers(db *dbCommon.DB) []httpCommon.Controller {
	responder := httpCommon.NewResponder()

	rafService := applicationCat.NewCategoryService(dbCat.NewCategoryRepository(db))

	rafReqHandler := NewCategoryRequestHandler()

	return []httpCommon.Controller{
		NewCategoryController(rafService, responder, rafReqHandler),
	}
}
