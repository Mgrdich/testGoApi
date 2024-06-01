package test_helpers

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MockMoviesRepository struct {
	GetAllFunc     func() ([]*models.Movie, error)
	GetByIDFunc    func(id uuid.UUID) (*models.Movie, error)
	SaveFunc       func(param models.CreateMovieParam) (*models.Movie, error)
	UpdateByIDFunc func(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	DeleteByIDFunc func(id uuid.UUID) error
}

func (m *MockMoviesRepository) GetAll(_ context.Context) ([]*models.Movie, error) {
	return m.GetAllFunc()
}

func (m *MockMoviesRepository) GetByID(_ context.Context, id uuid.UUID) (*models.Movie, error) {
	return m.GetByIDFunc(id)
}

func (m *MockMoviesRepository) Save(_ context.Context, param models.CreateMovieParam) (*models.Movie, error) {
	return m.SaveFunc(param)
}

func (m *MockMoviesRepository) UpdateByID(
	_ context.Context,
	id uuid.UUID,
	param models.UpdateMovieParam) (*models.Movie, error) {
	return m.UpdateByIDFunc(id, param)
}

func (m *MockMoviesRepository) DeleteByID(_ context.Context, id uuid.UUID) error {
	return m.DeleteByIDFunc(id)
}
