package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi.com/internal/controller"
	"testGoApi.com/internal/middlewares"
	"testGoApi.com/internal/services"
)

func GetMoviesRouter(r chi.Router) {
	movieStoreService := services.NewMoviesService()
	moviesController := controller.NewMoviesController(movieStoreService)
	r.Get("/", moviesController.HandleGetAllMovies)
	r.Post("/", moviesController.HandleCreateMovie)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.MovieCtx(movieStoreService.GetByID))
		r.Get("/", moviesController.HandleGetMovie)
		r.Put("/", moviesController.HandleUpdateMovie)
		r.Delete("/", moviesController.HandleDeleteMovie)
	})
}
