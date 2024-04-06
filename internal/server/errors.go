package server

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"` // application level error code
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var (
	ErrorNotFound            = &ErrorResponse{HTTPStatusCode: http.StatusNotFound, StatusText: "Resource not found."}
	ErrorBadRequest          = &ErrorResponse{HTTPStatusCode: http.StatusBadRequest, StatusText: "Bad request"}
	ErrorInternalServerError = &ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error",
	}
)

func ErrorConflict(err error) render.Renderer {
	return &ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		StatusText:     "Duplicate Key",
		ErrorText:      err.Error(),
	}
}
