package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testGoApi/internal/test_helpers"
)

var mockPersonService = test_helpers.NewMockPersonService()

func setPersonContext(person *models.Person) context.Context {
	return middlewares.SetPersonCtx(context.Background(), person)
}

func TestPersonController_HandleGetAllPerson(t *testing.T) {
	controller := NewPersonController(mockPersonService)

	req := test_helpers.NewRequest(t, http.MethodGet, "/person", nil)

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetAllPerson, nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response []personDTO

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	persons, err := mockPersonService.GetAll(context.Background())

	if err != nil {
		t.Errorf("Mock Person Servie Get all should not return err %v", err)
	}

	if len(response) != len(persons) {
		t.Errorf("Service Returned and the Response returned should match %v %v", len(response), len(persons))
	}
}

func TestPersonController_HandleGetPerson(t *testing.T) {
	controller := NewPersonController(mockPersonService)

	req := test_helpers.NewRequest(t, http.MethodGet, "/person/1", nil)
	person := &models.Person{ID: uuid.New(), FirstName: "testing", LastName: "Lastname"}
	ctx := setPersonContext(person)

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetPerson, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response personDTO

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.ID != person.ID {
		t.Errorf("ID does not match: %v %v", response.ID, person.ID)
	}

	if response.FirstName != person.FirstName {
		t.Errorf("Firstbname does not match: %v %v", response.FirstName, person.FirstName)
	}

	if response.LastName != person.LastName {
		t.Errorf("Lastname does not match: %v %v", response.LastName, person.LastName)
	}
}

func TestPersonController_HandleCreatePerson(t *testing.T) {
	personParams := CreatePersonRequest{
		FirstName: "Test",
		LastName:  "Test",
	}
	jsonData, err := json.Marshal(personParams)

	if err != nil {
		t.Errorf("Error encoding JSON: %v", err)
	}

	controller := NewPersonController(mockPersonService)

	req := test_helpers.NewRequest(t, http.MethodPost, "/person", bytes.NewBuffer(jsonData))

	rr := test_helpers.ExecuteRequest(req, controller.HandleCreatePerson, nil)

	if rr.Code != http.StatusCreated {
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusCreated, status)
		}
	}

	var response personDTO

	if err = json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.FirstName != personParams.FirstName {
		t.Errorf("Firstbname does not match: %v %v", response.FirstName, personParams.FirstName)
	}

	if response.LastName != personParams.LastName {
		t.Errorf("Lastname does not match: %v %v", response.LastName, personParams.LastName)
	}
}
