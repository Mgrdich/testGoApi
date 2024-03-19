package routes

import (
	".com/internal/controller"
	".com/internal/routes/middlewares"
	".com/internal/services"
	"github.com/go-chi/chi/v5"
)

func GetMoviesRouter(r chi.Router) {
	movieStoreService := services.NewMoviesService()
	moviesController := controller.NewMoviesController(movieStoreService)
	r.Get("/", moviesController.HandleGetAllMovies)
	r.Post("/", moviesController.HandleCreateMovie)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.MovieCtx(movieStoreService))
		r.Get("/", moviesController.HandleGetMovie)
		r.Put("/", moviesController.HandleUpdateMovie)
		r.Delete("/", moviesController.HandleDeleteMovie)
	})
}
