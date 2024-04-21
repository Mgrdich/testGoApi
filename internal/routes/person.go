package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi.com/internal/controller"
	"testGoApi.com/internal/middlewares"
	"testGoApi.com/internal/services"
)

func GetPersonRouter(r chi.Router) {
	personStoreService := services.NewPersonService()
	personController := controller.NewPersonController(personStoreService)

	r.Get("/", personController.HandleGetAllPerson)
	r.Post("/", personController.HandleCreatePerson)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(middlewares.PersonCtx(personStoreService.GetByID))
		r.Get("/", personController.HandleGetPerson)
	})
}
