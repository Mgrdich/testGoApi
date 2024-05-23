package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/services"
)

func GetMoviesRouter(movieStoreService services.MovieService) func(router chi.Router) {
	return func(r chi.Router) {
		moviesController := controller.NewMoviesController(movieStoreService)
		r.Get("/", moviesController.HandleGetAllMovies)
		r.Post("/", moviesController.HandleCreateMovie)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.GetContextIdFunc(movieStoreService.Get, middlewares.SetMovieCtx))
			r.Get("/", moviesController.HandleGetMovie)
			r.Put("/", moviesController.HandleUpdateMovie)
			r.Delete("/", moviesController.HandleDeleteMovie)
		})
	}
}
