package middlewares

import (
	"context"

	"testGoApi/internal/models"
)

type movieContentKeyType int

const movieContextKey movieContentKeyType = 0

// GetMovieCtx GetMovieContext retrieves movie information from the context
func GetMovieCtx(ctx context.Context) (*models.Movie, bool) {
	movie, ok := ctx.Value(movieContextKey).(*models.Movie)
	return movie, ok
}

// SetMovieCtx SetMovieContext sets movie information in the context
func SetMovieCtx(ctx context.Context, movie *models.Movie) context.Context {
	return context.WithValue(ctx, movieContextKey, movie)
}
