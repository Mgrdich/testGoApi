package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"testGoApi/internal/controller"
	"testGoApi/internal/server"
	"testGoApi/internal/services"
)

type ApplicationServices struct {
	MovieService  services.MovieService
	PersonService services.PersonService
}

func AddRoutes(s *server.Server, services *ApplicationServices) {
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))

	s.Router.Get("/health", controller.HandleGetHealth)

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/movies", GetMoviesRouter(services.MovieService))
		r.Route("/person", GetPersonRouter(services.PersonService))
	})

	s.Router.Mount("/swagger", httpSwagger.WrapHandler)
}
