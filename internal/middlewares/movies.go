package middlewares

import (
	"context"
	"net/http"

	".com/internal/db"
	".com/internal/models"
	".com/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// Define a custom type for the context key
type movieContextKeyType struct{}

// Implement the String method for the context key type
func (k *movieContextKeyType) String() string {
	return "movie"
}

// Create a global instance of the movie context key
var movieContextKey = &movieContextKeyType{}

// GetMovieContext retrieves movie information from the context
func GetMovieCtx(ctx context.Context) *models.Movie {
	return ctx.Value(movieContextKey).(*models.Movie)
}

// SetMovieContext sets movie information in the context
func SetMovieCtx(ctx context.Context, movie *models.Movie) context.Context {
	return context.WithValue(ctx, movieContextKey, movie)
}

// MovieContextMiddleware adds movie information to the request context
func MovieCtx(moviesStore db.MoviesStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			movieID := chi.URLParam(r, "id")
			if movieID == "" {
				render.Render(w, r, server.ErrorBadRequest)
				return
			}

			id, err := uuid.Parse(movieID)
			if err != nil {
				render.Render(w, r, server.ErrorBadRequest)
				return
			}

			movie, err := moviesStore.GetByID(id)

			if err != nil {
				render.Render(w, r, server.ErrorNotFound)
				return
			}

			// Set movie information in the request context
			ctx := SetMovieCtx(r.Context(), movie)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
