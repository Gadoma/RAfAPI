package http

import "net/http"

type CategoryResponder interface {
	RespondSuccessOk(w http.ResponseWriter, data any, n int)
	RespondSuccessNoContent(w http.ResponseWriter)
	RespondErrorNotFound(w http.ResponseWriter)
	RespondError(w http.ResponseWriter, message string, httpStatus int)
}
