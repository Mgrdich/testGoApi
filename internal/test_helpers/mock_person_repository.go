package test_helpers

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MockPersonRepository struct {
	GetAllFunc  func() ([]*models.Person, error)
	GetByIDFunc func(id uuid.UUID) (*models.Person, error)
	SaveFunc    func(param models.CreatePerson) (*models.Person, error)
}

func (m *MockPersonRepository) GetAll(_ context.Context) ([]*models.Person, error) {
	return m.GetAllFunc()
}

func (m *MockPersonRepository) GetByID(_ context.Context, id uuid.UUID) (*models.Person, error) {
	return m.GetByIDFunc(id)
}

func (m *MockPersonRepository) Save(_ context.Context, param models.CreatePerson) (*models.Person, error) {
	return m.SaveFunc(param)
}
