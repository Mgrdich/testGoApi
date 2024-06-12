package server

import (
	"net/http"

	"github.com/go-chi/render"
)

// HTTPError represents the HTTP error details.
type HTTPError struct {
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
}

// ErrorResponse represents the response model for errors.
type ErrorResponse struct {
	*HTTPError
	Err       error  `json:"-"`
	AppCode   int64  `json:"code,omitempty"` // application level error code
	ErrorText string `json:"error,omitempty"`
}

// Render renders the error response.
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrorRouteNotFound represents a route not found error.
var ErrorRouteNotFound = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Error page not found",
	},
}

// ErrorMethodNotAllowed represents a Method Not allowed error.
var ErrorMethodNotAllowed = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Method Not Allowed",
	},
}

// ErrorNotFound represents a not found error.
var ErrorNotFound = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     "Resource not found.",
	},
}

// ErrorBadRequest represents a bad request error.
var ErrorBadRequest = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad request",
	},
}

// ErrorInternalServerError represents an internal server error.
var ErrorInternalServerError = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error",
	},
}

// ErrorForbidden represents an internal server error.
var ErrorForbidden = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusForbidden,
		StatusText:     "Forbidden",
	},
}

var ErrorUnauthorized = &ErrorResponse{
	HTTPError: &HTTPError{
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     "Unauthorized",
	},
}

// ErrorConflict returns an error response for a conflict error.
func ErrorConflict(err error) render.Renderer {
	return &ErrorResponse{
		HTTPError: &HTTPError{
			HTTPStatusCode: http.StatusConflict,
			StatusText:     "Duplicate Key",
		},
		Err:       err,
		ErrorText: err.Error(),
	}
}
