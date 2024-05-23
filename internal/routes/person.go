package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/services"
)

func GetPersonRouter(personStoreService services.PersonService) func(router chi.Router) {
	return func(r chi.Router) {
		personController := controller.NewPersonController(personStoreService)

		r.Get("/", personController.HandleGetAllPerson)
		r.Post("/", personController.HandleCreatePerson)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.GetContextIdFunc(personStoreService.Get, middlewares.SetPersonCtx))
			r.Get("/", personController.HandleGetPerson)
		})
	}
}
