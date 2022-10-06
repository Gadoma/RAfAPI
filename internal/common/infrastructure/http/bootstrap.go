package http

import dbCommon "github.com/gadoma/rafapi/internal/common/infrastructure/database"

type Bootstrapper interface {
	BootstrapControllers(db *dbCommon.DB) []Controller
}
