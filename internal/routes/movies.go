package routes

import (
	".com/internal/controller"
	".com/internal/services"
	"github.com/go-chi/chi/v5"
)

func GetMoviesRouter(r chi.Router) {

	moviesController := controller.NewMoviesController(services.NewMoviesService())

	r.Get("/", moviesController.HandleGetAllMovies)
	r.Post("/", moviesController.HandleCreateMovie)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", moviesController.HandleGetMovie)
		r.Put("/", moviesController.HandleUpdateMovie)
		r.Delete("/", moviesController.HandleDeleteMovie)
	})
}
