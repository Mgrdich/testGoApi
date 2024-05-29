package routes

import (
	"github.com/go-chi/chi/v5"
	"testGoApi/internal/controller"
	"testGoApi/internal/services"
)

func GetUserRouter(userService services.UserService, tokenService services.TokenService) func(router chi.Router) {
	return func(r chi.Router) {
		userController := controller.NewUserController(userService, tokenService)

		r.Post("/register", userController.HandleRegisterUser)
		r.Post("/login", userController.HandleLoginUser)
		r.Get("/me", userController.HandleUserMe) // TODO should have protection
	}
}
