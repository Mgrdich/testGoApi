package controller

import (
	"net/http"

	"github.com/go-chi/render"
	"testGoApi/internal/server"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, server.ErrorRouteNotFound)
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, server.ErrorMethodNotAllowed)
}
