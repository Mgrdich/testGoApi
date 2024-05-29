package repository

import (
	"context"
	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MovieRepository interface {
	GetAll(ctx context.Context) ([]*models.Movie, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Movie, error)
	Save(ctx context.Context, param models.CreateMovieParam) (*models.Movie, error)
	UpdateByID(ctx context.Context, id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type PersonRepository interface {
	GetAll() ([]*models.Person, error)
	GetByID(id uuid.UUID) (*models.Person, error)
	Save(param models.CreatePerson) (*models.Person, error)
}
