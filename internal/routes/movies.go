package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/services"
)

func GetMoviesRouter(movieService services.MovieService) func(router chi.Router) {
	return func(r chi.Router) {
		moviesController := controller.NewMoviesController(movieService)

		r.MethodNotAllowed(controller.HandleMethodNotAllowed)
		r.Get("/", moviesController.HandleGetAllMovies)
		r.Post("/", moviesController.HandleCreateMovie)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.GetContextIdFunc(movieService.Get, middlewares.SetMovieCtx))
			r.Use(
				middlewares.AllowedMethods(
					http.MethodGet,
					http.MethodPut,
					http.MethodDelete,
				),
				middlewares.GetContextIdFunc(movieStoreService.Get, middlewares.SetMovieCtx),
			)
			r.Get("/", moviesController.HandleGetMovie)
			r.Put("/", moviesController.HandleUpdateMovie)
			r.Delete("/", moviesController.HandleDeleteMovie)
		})
	}
}
