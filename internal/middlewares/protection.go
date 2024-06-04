package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"testGoApi/internal/models"
	"testGoApi/internal/server"
	"testGoApi/internal/services"
)

type userContextKeyType int

const userContextKey userContextKeyType = 0

// GetTokenizedUserCtx retrieves person information from the context
func GetTokenizedUserCtx(ctx context.Context) (*models.TokenizedUser, bool) {
	person, ok := ctx.Value(userContextKey).(*models.TokenizedUser)
	return person, ok
}

// setTokenizedUserCtx sets person information in the context
func setTokenizedUserCtx(ctx context.Context, user *models.TokenizedUser) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// Authentication is simple middleware to deal with
func Authentication(tokenService services.TokenService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				log.Println("Missing authorization header")

				_ = render.Render(w, r, server.ErrorUnauthorized)

				return
			}

			tokenString = tokenString[len("Bearer "):]

			if err := tokenService.VerifyJWT(tokenString); err != nil {
				_ = render.Render(w, r, server.ErrorForbidden)
				return
			}

			user, err := tokenService.ParseJWT(tokenString)

			if err != nil {
				_ = render.Render(w, r, server.ErrorForbidden)
				return
			}

			ctx := setTokenizedUserCtx(r.Context(), user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Authorized is simple middleware to deal with authorization of a particular role
// Bearing in mind this needs Authentication to work
func Authorized(roles ...models.UserRole) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := GetTokenizedUserCtx(r.Context())

			if !ok {
				_ = render.Render(w, r, server.ErrorForbidden)
				return
			}

			isAuthorized := false

			for _, role := range roles {
				if userRole, ok := models.LookUpRole(user.Role); role == userRole && ok {
					isAuthorized = true
					break
				}
			}

			if !isAuthorized {
				_ = render.Render(w, r, server.ErrorForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
