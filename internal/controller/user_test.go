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

const username = "username"
const password = "Password123@"

func TestUserController_HandleLoginUser(t *testing.T) {
	controller := NewUserController(test_helpers.NewMockUserService())
	param := loginUserRequest{
		Username: username,
		Password: password,
	}
	jsonData, err := json.Marshal(param)

	if err != nil {
		t.Errorf("Error encoding JSON: %v", err)
	}

	req := test_helpers.NewRequest(t, http.MethodPost, "/login", bytes.NewBuffer(jsonData))

	rr := test_helpers.ExecuteRequest(req, controller.HandleLoginUser, nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response tokenDTO

	if err = json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if len(response.Token) == 0 {
		t.Error("token should not be empty")
	}
}

func TestUserController_HandleRegisterUser(t *testing.T) {
	controller := NewUserController(test_helpers.NewMockUserService())
	param := registerUserRequest{
		Username:       username,
		Password:       password,
		RepeatPassword: password,
	}
	jsonData, err := json.Marshal(param)

	if err != nil {
		t.Errorf("Error encoding JSON: %v", err)
	}

	req := test_helpers.NewRequest(t, http.MethodPost, "/register", bytes.NewBuffer(jsonData))

	rr := test_helpers.ExecuteRequest(req, controller.HandleRegisterUser, nil)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response OKDto

	if err = json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.OK != true {
		t.Error("Response back does not match the intention")
	}
}

func TestUserController_HandleUserMe(t *testing.T) {
	user := &models.TokenizedUser{
		ID:       uuid.New(),
		Username: username,
		Role:     models.LookUpRoleString(models.BasicRole),
	}
	ctx := middlewares.SetTokenizedUserCtx(context.Background(), user)
	controller := NewUserController(test_helpers.NewMockUserService())

	req := test_helpers.NewRequest(t, http.MethodGet, "/me", nil)

	rr := test_helpers.ExecuteRequest(req, controller.HandleUserMe, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response userDto

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.ID != user.ID {
		t.Errorf("ID does not match: %v %v", response.ID, user.ID)
	}

	if response.Username != user.Username {
		t.Errorf("Username does not match: %v %v", response.Username, user.Username)
	}
}
