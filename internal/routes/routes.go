package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"testGoApi/internal/controller"
	"testGoApi/internal/db"
	"testGoApi/internal/server"
)

type ApplicationServices struct {
	MovieStore  db.MoviesStore
	PersonStore db.PersonStore
}

func AddRoutes(s *server.Server, services *ApplicationServices) {
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))

	s.Router.Get("/health", controller.HandleGetHealth)

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/movies", GetMoviesRouter(services.MovieStore))
		r.Route("/person", GetPersonRouter(services.PersonStore))
	})
}
