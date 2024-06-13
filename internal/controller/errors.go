package controller

import (
	"net/http"

	"github.com/go-chi/render"
	"testGoApi/internal/server"
)

// HandleNotFound handles the 404 errors.
// @Summary Not Found
// @Description This endpoint handles requests to routes that do not exist.
// @Tags errors
// @Produce json
// @Failure 404 {object} server.HTTPError
// @Router /example-notfound [get]
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, server.ErrorRouteNotFound)
}

// HandleMethodNotAllowed handles the 405 errors.
// @Summary Method Not Allowed
// @Description This endpoint handles requests with HTTP methods that are not allowed on the given route.
// @Tags errors
// @Produce json
// @Failure 405 {object} server.HTTPError
// @Router /api/v1/person [patch]
func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, server.ErrorMethodNotAllowed)
}
