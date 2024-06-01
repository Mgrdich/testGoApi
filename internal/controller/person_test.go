package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testGoApi/internal/test_helpers"
	"testing"
)

func setPersonContext() context.Context {
	return middlewares.SetPersonCtx(context.Background(), &models.Person{})
}

func TestPersonController_HandleGetAllPerson(t *testing.T) {
	controller := NewPersonController(&test_helpers.MockPersonService{})

	req := test_helpers.NewRequest(t, http.MethodGet, "/person", nil)

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetAllPerson, nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestPersonController_HandleGetPerson(t *testing.T) {

	controller := NewPersonController(&test_helpers.MockPersonService{})

	req := test_helpers.NewRequest(t, http.MethodGet, "/person/1", nil)
	ctx := setPersonContext()

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetPerson, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestPersonController_HandleCreatePerson(t *testing.T) {
	personParams := CreatePersonRequest{
		FirstName: "Test",
		LastName:  "Test",
	}
	jsonData, _ := json.Marshal(personParams)

	controller := NewPersonController(&test_helpers.MockPersonService{})

	req := test_helpers.NewRequest(t, http.MethodPost, "/person", bytes.NewBuffer(jsonData))

	rr := test_helpers.ExecuteRequest(req, controller.HandleCreatePerson, nil)

	if rr.Code != http.StatusCreated {
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusCreated, status)
		}
	}
}
