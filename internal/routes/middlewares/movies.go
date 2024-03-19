package middlewares

import (
	"context"
	"fmt"
	"net/http"

	".com/internal/db"
	".com/internal/models"
	".com/internal/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type movieCtxKeyType string

const movieCtxKey movieCtxKeyType = "movie"

// Define other keys as needed

func GetMovieCtx(ctx context.Context) models.Movie {
	return ctx.Value(movieCtxKey).(models.Movie)
}

func SetMovieCtx(ctx context.Context, movie *models.Movie) context.Context {
	return context.WithValue(ctx, movieCtxKey, movie)
}

// MovieCtx middleware adds movie information to the request context
func MovieCtx(moviesStore db.MoviesStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println(movie)
			if err != nil {
				render.Render(w, r, server.ErrorNotFound)
				return
			}

			// Set movie information in the request context

			ctx := SetMovieCtx(r.Context(), movie)

			fmt.Println(ctx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
