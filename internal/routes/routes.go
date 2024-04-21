package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"testGoApi.com/internal/controller"
	"testGoApi.com/internal/server"
)

func AddRoutes(s *server.Server) {
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))

	s.Router.Get("/health", controller.HandleGetHealth)

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/movies", GetMoviesRouter)
		r.Route("/person", GetPersonRouter)
	})
}
