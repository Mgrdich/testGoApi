package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testGoApi/internal/test_helpers"
)

var mockMoviesService = test_helpers.NewMockMovieService()

func setMovieCtx() context.Context {
	return middlewares.SetMovieCtx(context.Background(), &models.Movie{})
}

func TestMoviesController_HandleGetAllMovies(t *testing.T) {
	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodGet, "/movies", nil)

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetAllMovies, nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestMoviesController_HandleGetMovie(t *testing.T) {
	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodGet, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetMovie, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestMoviesController_HandleCreateMovie(t *testing.T) {
	createParam := CreateMovieRequest{
		Title:       "title",
		Director:    "Director",
		TicketPrice: 2,
	}

	jsonData, err := json.Marshal(createParam)

	if err != nil {
		t.Errorf("Error encoding JSON: %v", err)
	}

	controller := NewMoviesController(mockMoviesService)

	// json content-type
	req := test_helpers.NewRequest(t, http.MethodPost, "/movies", bytes.NewBuffer(jsonData))

	rr := test_helpers.ExecuteRequest(req, controller.HandleCreateMovie, nil)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusCreated, status)
	}
}

func TestMoviesController_HandleUpdateMovie(t *testing.T) {
	updateParam := UpdateMovieRequest{
		Title:       "title",
		Director:    "Director",
		TicketPrice: 2,
	}

	jsonData, err := json.Marshal(updateParam)

	if err != nil {
		t.Errorf("Error encoding JSON: %v", err)
	}

	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodPut, "/movies/1", bytes.NewBuffer(jsonData))
	ctx := setMovieCtx()

	rr := test_helpers.ExecuteRequest(req, controller.HandleUpdateMovie, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestMoviesController_HandleDeleteMovie(t *testing.T) {
	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodDelete, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := test_helpers.ExecuteRequest(req, controller.HandleDeleteMovie, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}
