package routes

import (
	".com/internal/controller"
	".com/internal/services"
	"github.com/go-chi/chi/v5"
)

func GetPersonRouter(r chi.Router) {
	personStoreService := services.NewPersonService()
	personController := controller.NewPersonController(personStoreService)

	r.Get("/", personController.HandleGetAllPerson)
	r.Post("/", personController.HandleCreatePerson)

	r.Route("/{id}", func(r chi.Router) {
		//r.Use(middlewares.PersonCtx(personStoreService))
		r.Get("/", personController.HandleGetPerson)
	})
}
