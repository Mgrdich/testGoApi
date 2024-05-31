package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testing"
	"time"
)

type MockPersonService struct{}

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

func setPersonContext() context.Context {
	return middlewares.SetPersonCtx(context.Background(), &models.Person{})
}

func TestPersonController_HandleGetAllPerson(t *testing.T) {
	controller := NewPersonController(&MockPersonService{})

	req := NewRequest(t, http.MethodGet, "/person", nil)

	rr := ExecuteRequest(req, controller.HandleGetAllPerson, nil)

	CheckStatusOK(t, rr)
}

func TestPersonController_HandleGetPerson(t *testing.T) {

	controller := NewPersonController(&MockPersonService{})

	req := NewRequest(t, http.MethodGet, "/person/1", nil)
	ctx := setPersonContext()

	rr := ExecuteRequest(req, controller.HandleGetPerson, ctx)

	CheckStatusOK(t, rr)
}

func TestPersonController_HandleCreatePerson(t *testing.T) {
	personParams := CreatePersonRequest{
		FirstName: "Test",
		LastName:  "Test",
	}
	jsonData, _ := json.Marshal(personParams)

	controller := NewPersonController(&MockPersonService{})

	req := NewRequest(t, http.MethodPost, "/person", bytes.NewBuffer(jsonData))

	rr := ExecuteRequest(req, controller.HandleCreatePerson, nil)

	CheckStatusCreated(t, rr)
}
