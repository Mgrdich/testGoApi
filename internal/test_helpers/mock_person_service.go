package test_helpers

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MockPersonService struct{}

func NewMockPersonService() *MockPersonService {
	return &MockPersonService{}
}

var person = &models.Person{
	ID:        uuid.New(),
	FirstName: "Test",
	LastName:  "Test",
	CreatedAt: time.Now(),
}

func (m *MockPersonService) GetAll(_ context.Context) ([]*models.Person, error) {
	return []*models.Person{
		person,
	}, nil
}
func (m *MockPersonService) Get(_ context.Context, id uuid.UUID) (*models.Person, error) {
	if id != person.ID {
		return nil, errors.New("person not found")
	}

	return person, nil
}
func (m *MockPersonService) Create(_ context.Context, param models.CreatePerson) (*models.Person, error) {
	return &models.Person{
		ID:        uuid.New(),
		FirstName: param.FirstName,
		LastName:  param.LastName,
		CreatedAt: time.Now(),
	}, nil
}
