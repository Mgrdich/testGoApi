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
// @Success 404 {object} render.Renderer "Route not found"
func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, server.ErrorRouteNotFound)
}

// HandleMethodNotAllowed handles the 405 errors.
// @Summary Method Not Allowed
// @Description This endpoint handles requests with HTTP methods that are not allowed on the given route.
// @Tags errors
// @Produce json
// @Success 405 {object} render.Renderer "Method not allowed"
func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, server.ErrorMethodNotAllowed)
}
