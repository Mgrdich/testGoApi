package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/db"
	"testGoApi/internal/middlewares"
)

func GetMoviesRouter(movieStoreService db.MoviesStore) func(router chi.Router) {
	return func(r chi.Router) {
		moviesController := controller.NewMoviesController(movieStoreService)
		r.Get("/", moviesController.HandleGetAllMovies)
		r.Post("/", moviesController.HandleCreateMovie)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.GetContextIdFunc(movieStoreService.GetByID, middlewares.SetMovieCtx))
			r.Get("/", moviesController.HandleGetMovie)
			r.Put("/", moviesController.HandleUpdateMovie)
			r.Delete("/", moviesController.HandleDeleteMovie)
		})
	}
}
