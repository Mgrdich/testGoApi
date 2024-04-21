package middlewares

import (
	"context"
	"net/http"

	".com/internal/models"
	".com/internal/util"
)

const movieContextKey = "middlewares.movieContextKey"

// GetMovieCtx GetMovieContext retrieves movie information from the context
func GetMovieCtx(ctx context.Context) (*models.Movie, bool) {
	movie, ok := ctx.Value(movieContextKey).(*models.Movie)
	return movie, ok
}

// SetMovieCtx SetMovieContext sets movie information in the context
func SetMovieCtx(ctx context.Context, movie *models.Movie) context.Context {
	return context.WithValue(ctx, movieContextKey, movie)
}

// MovieCtx Middleware adds person information to the request context
func MovieCtx(getIdFunc util.GetByIDFunc[models.Movie]) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			person, err := CheckSlugId(w, r, getIdFunc)
			if err != nil {
				return
			}
			// Set person information in the request context
			ctx := SetMovieCtx(r.Context(), person)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
