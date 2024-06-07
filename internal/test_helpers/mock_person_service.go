package test_helpers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MockPersonService struct{}

func NewMockPersonService() *MockPersonService {
	return &MockPersonService{}
}

func (m *MockPersonService) GetAll(_ context.Context) ([]*models.Person, error) {
	persons := []*models.Person{
		{ID: uuid.New()},
		{ID: uuid.New()},
		{ID: uuid.New()},
	}

	return persons, nil
}

func (m *MockPersonService) Get(_ context.Context, id uuid.UUID) (*models.Person, error) {
	return &models.Person{ID: id}, nil
}

func (m *MockPersonService) Create(_ context.Context, param models.CreatePerson) (*models.Person, error) {
	return &models.Person{
		ID:        uuid.New(),
		FirstName: param.FirstName,
		LastName:  param.LastName,
		CreatedAt: time.Now(),
	}, nil
}
