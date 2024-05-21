package repository

import (
	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MovieRepositoryImpl struct {
}

func NewMovieRepositoryImpl() *MovieRepositoryImpl {
	return &MovieRepositoryImpl{}
}

func (m *MovieRepositoryImpl) GetAll() ([]models.Movie, error) {
	return nil, nil
}

func (m *MovieRepositoryImpl) GetByID() (*models.Movie, error) {
	return nil, nil
}

func (m *MovieRepositoryImpl) Create(param models.CreateMovieParam) (*models.Movie, error) {
	return nil, nil
}

func (m *MovieRepositoryImpl) UpdateById(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	return nil, nil
}

func (m *MovieRepositoryImpl) DeleteById(id uuid.UUID) (*models.Movie, error) {
	return nil, nil
}
