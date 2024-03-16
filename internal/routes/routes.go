package routes

import (
	".com/internal/controller"
	".com/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func AddRoutes(s *server.Server) {
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))

	s.Router.Get("/health", controller.HandleGetHealth)

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/movies", GetMoviesRouter)
		r.Route("/person", GetPersonRouter)
	})
}
