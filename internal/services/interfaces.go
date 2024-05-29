package services

import (
	"context"
	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MovieService interface {
	GetAll(ctx context.Context) ([]*models.Movie, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Movie, error)
	Create(ctx context.Context, param models.CreateMovieParam) (*models.Movie, error)
	Update(ctx context.Context, id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type PersonService interface {
	GetAll() ([]*models.Person, error)
	Get(id uuid.UUID) (*models.Person, error)
	Create(param models.CreatePerson) (*models.Person, error)
}
