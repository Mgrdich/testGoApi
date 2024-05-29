package middlewares

import (
	"net/http"

	"testGoApi/internal/models"
)

// Authentication is simple middleware to deal with
func Authentication(verifyToken func()) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}

// Authorized is simple middleware to deal with authorization of a particular role
// Bearing in mind this needs Authentication to work
func Authorized(role models.UserRole) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}
