package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"testGoApi/configs"
	"testGoApi/internal/controller"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/server"
	"testGoApi/internal/services"
	"testGoApi/internal/util"
)

type ApplicationServices struct {
	MovieService  services.MovieService
	PersonService services.PersonService
	UserService   services.UserService
	TokenService  services.TokenService
}

func AddRoutes(s *server.Server, services *ApplicationServices) {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(render.SetContentType(render.ContentTypeJSON))

	s.Router.Get("/health", controller.HandleGetHealth)

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/user", GetUserRouter(services.UserService))

		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.Authentication(services.TokenService))
			r.Route("/movies", GetMoviesRouter(services.MovieService))
			r.Route("/person", GetPersonRouter(services.PersonService))
		})
	})

	if configs.GetAppConfig().Environment == util.DevEnvironment {
		s.Router.Mount("/swagger", httpSwagger.WrapHandler)
	}
}
