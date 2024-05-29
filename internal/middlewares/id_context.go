package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"testGoApi/internal/models"
	"testGoApi/internal/server"
	"testGoApi/internal/util"
)

func GetContextIdFunc[K models.Models](getByIdFunc util.GetByIDFunc[K],
	contextSetter func(ctx context.Context, value *K) context.Context) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			movieID := chi.URLParam(r, "id")
			if movieID == "" {
				_ = render.Render(w, r, server.ErrorBadRequest)
				return
			}

			id, err := uuid.Parse(movieID)

			if err != nil {
				_ = render.Render(w, r, server.ErrorBadRequest)
				return
			}

			storeValue, err := getByIdFunc(r.Context(), id)

			if err != nil {
				_ = render.Render(w, r, server.ErrorNotFound)
				return
			}

			// Set movie information in the request context
			ctx := contextSetter(r.Context(), storeValue)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
