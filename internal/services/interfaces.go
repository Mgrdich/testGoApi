package services

import (
	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MovieService interface {
	GetAll() ([]*models.Movie, error)
	GetByID(id uuid.UUID) (*models.Movie, error)
	Create(param models.CreateMovieParam) (*models.Movie, error)
	Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	Delete(id uuid.UUID) error
}

type PersonService interface {
	GetAll() ([]*models.Person, error)
	GetByID(id uuid.UUID) (*models.Person, error)
	Create(param models.CreatePerson) (*models.Person, error)
}
