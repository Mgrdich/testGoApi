package middlewares

import (
	"context"
	"errors"
	"net/http"

	".com/internal/db"
	".com/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func CheckSlugId[K any](w http.ResponseWriter, r *http.Request, store db.GetByIdStore[K]) (*K, error) {
	slugId := chi.URLParam(r, "id")
	if slugId == "" {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return nil, errors.New("slug ID is empty")
	}

	id, err := uuid.Parse(slugId)
	if err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return nil, errors.New("invalid slug ID")
	}

	storeValue, err := store.GetByID(id)
	if err != nil {
		_ = render.Render(w, r, server.ErrorNotFound)
		return nil, errors.New("ID not found in store")
	}

	return storeValue, nil
}

func GetContextIdFunc[K any](store db.GetByIdStore[K],
	contextSetter func(ctx context.Context, value *K) context.Context) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
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

			storeValue, err := store.GetByID(id)

			if err != nil {
				_ = render.Render(w, r, server.ErrorNotFound)
				return
			}

			// Set movie information in the request context
			ctx := contextSetter(r.Context(), storeValue)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
