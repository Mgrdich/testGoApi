package middlewares

import (
	"context"

	"testGoApi/internal/models"
)

type personContextKeyType int

const personContextKey personContextKeyType = 0

// GetPersonCtx retrieves person information from the context
func GetPersonCtx(ctx context.Context) (*models.Person, bool) {
	person, ok := ctx.Value(personContextKey).(*models.Person)
	return person, ok
}

// SetPersonCtx sets person information in the context
func SetPersonCtx(ctx context.Context, person *models.Person) context.Context {
	return context.WithValue(ctx, personContextKey, person)
}
