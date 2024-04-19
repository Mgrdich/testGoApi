package middlewares

import (
	"context"
	"net/http"

	".com/internal/db"
	".com/internal/models"
)

const personContextKey = "middlewares.personContextKey"

// GetPersonCtx retrieves person information from the context
func GetPersonCtx(ctx context.Context) (*models.Person, bool) {
	person, ok := ctx.Value(personContextKey).(*models.Person)
	return person, ok
}

// SetPersonCtx sets person information in the context
func SetPersonCtx(ctx context.Context, person *models.Person) context.Context {
	return context.WithValue(ctx, personContextKey, person)
}

// PersonCtxMiddleware adds person information to the request context
func PersonCtx(personStore db.PersonStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			person, err := CheckSlugId[models.Person](w, r, personStore)
			if err != nil {
				return
			}
			// Set person information in the request context
			ctx := SetPersonCtx(r.Context(), person)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
