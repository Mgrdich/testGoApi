package middlewares

import (
	"context"
	"log"
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
			slugId := chi.URLParam(r, "id")
			if slugId == "" {
				log.Println(util.NewSlugIDIsNotDefined())

				_ = render.Render(w, r, server.ErrorBadRequest)

				return
			}

			id, err := uuid.Parse(slugId)

			if err != nil {
				log.Println(err)

				_ = render.Render(w, r, server.ErrorBadRequest)

				return
			}

			storeValue, err := getByIdFunc(r.Context(), id)

			if err != nil {
				log.Println(err)

				_ = render.Render(w, r, server.ErrorNotFound)

				return
			}

			// Set data information in the request context
			ctx := contextSetter(r.Context(), storeValue)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AllowedMethods(methods ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !util.Includes(methods, r.Method) {
				_ = render.Render(w, r, server.ErrorMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func NotAllowedMethods(methods ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if util.Includes(methods, r.Method) {
				_ = render.Render(w, r, server.ErrorMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
