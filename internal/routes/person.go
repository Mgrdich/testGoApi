package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testGoApi/internal/services"
)

func GetPersonRouter(personService services.PersonService) func(router chi.Router) {
	return func(r chi.Router) {
		personController := controller.NewPersonController(personService)

		r.With(middlewares.Authorized(models.AdminRole)).Get("/", personController.HandleGetAllPerson)
		r.Post("/", personController.HandleCreatePerson)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(middlewares.GetContextIdFunc(personService.Get, middlewares.SetPersonCtx))
			r.Get("/", personController.HandleGetPerson)
		})
	}
}
