package services

import (
	"context"
	"errors"
	"testing"
	"time"

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

func TestPersonService_GetAll(t *testing.T) {
	expectedPersons := []*models.Person{
		{ID: uuid.New(), FirstName: "Test 1", LastName: "Test 1"},
		{ID: uuid.New(), FirstName: "Test 2", LastName: "Test 2"},
	}

	GetAllFunc := func() ([]*models.Person, error) {
		return expectedPersons, nil
	}
	mockRepo := &MockPersonRepository{
		GetAllFunc: GetAllFunc,
	}

	personService := NewPersonServiceImpl(mockRepo)

	persons, err := personService.GetAll(context.Background())

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(persons) != len(expectedPersons) {
		if len(persons) != len(expectedPersons) {
			t.Fatalf("expected %d movies, got %d", len(expectedPersons), len(persons))
		}
	}
}

func TestPersonService_Get(t *testing.T) {
	id := uuid.New()
	expectedPerson := &models.Person{
		ID:        id,
		FirstName: "Test 1",
		LastName:  "Test 1",
		CreatedAt: time.Now(),
	}

	getByIDFunc := func(personId uuid.UUID) (*models.Person, error) {
		if id == personId {
			return expectedPerson, nil
		}

		return nil, errors.New("person not found")
	}

	mockRepo := &MockPersonRepository{
		GetByIDFunc: getByIDFunc,
	}

	personService := NewPersonServiceImpl(mockRepo)

	person, err := personService.Get(context.Background(), id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if person != expectedPerson {
		t.Errorf("expected person to be %v, got %v", expectedPerson, person)
	}
}

func TestPersonService_Create(t *testing.T) {
	createParams := models.CreatePerson{
		FirstName: "Test 1",
		LastName:  "Test 1",
	}
	expectedPerson := &models.Person{
		ID:        uuid.New(),
		FirstName: createParams.FirstName,
		LastName:  createParams.LastName,
		CreatedAt: time.Now(),
	}

	SaveFunc := func(param models.CreatePerson) (*models.Person, error) {
		if param.FirstName != "" && param.LastName != "" {
			return expectedPerson, nil
		}

		return nil, errors.New("first name and last name must be provided")
	}
	mockRepo := &MockPersonRepository{
		SaveFunc: SaveFunc,
	}

	personService := NewPersonServiceImpl(mockRepo)

	person, err := personService.Create(context.Background(), createParams)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if person != expectedPerson {
		t.Fatalf("expected movie %v, got %v", expectedPerson, person)
	}
}
