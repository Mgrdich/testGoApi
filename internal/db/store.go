package db

import (
	".com/internal/models"
	"github.com/google/uuid"
)

type MoviesStore interface {
	GetAll() ([]*models.Movie, error)
	GetByID(id uuid.UUID) (*models.Movie, error)
	Create(param models.CreateMovieParam) (*models.Movie, error)
	Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	Delete(id uuid.UUID) error
}

type PersonStore interface {
	GetAll() ([]*models.Person, error)
	GetByID(id uuid.UUID) (*models.Person, error)
	Create(param models.CreatePerson) (*models.Person, error)
}

type GetByIdStore[K any] interface {
	GetByID(id uuid.UUID) (*K, error)
}
