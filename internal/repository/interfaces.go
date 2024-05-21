package repository

import (
	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MovieRepository interface {
	GetAll() ([]*models.Movie, error)
	GetByID(id uuid.UUID) (*models.Movie, error)
	Save(param models.CreateMovieParam) (*models.Movie, error)
	UpdateByID(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	DeleteByID(id uuid.UUID) error
}

type PersonRepository interface {
	GetAll() ([]*models.Person, error)
	GetByID(id uuid.UUID) (*models.Person, error)
	Save(param models.CreatePerson) (*models.Person, error)
}
