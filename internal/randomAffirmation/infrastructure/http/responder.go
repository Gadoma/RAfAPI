package http

import "net/http"

type RandomAffirmationResponder interface {
	RespondSuccessOk(w http.ResponseWriter, data any, n int)
	RespondError(w http.ResponseWriter, message string, httpStatus int)
}
