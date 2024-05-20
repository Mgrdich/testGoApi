package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/db"
	"testGoApi/internal/middlewares"
)

func GetPersonRouter(personStoreService db.PersonStore) func(router chi.Router) {
	return func(r chi.Router) {
		personController := controller.NewPersonController(personStoreService)

		r.Get("/", personController.HandleGetAllPerson)
		r.Post("/", personController.HandleCreatePerson)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.GetContextIdFunc(personStoreService.GetByID, middlewares.SetPersonCtx))
			r.Get("/", personController.HandleGetPerson)
		})
	}
}
