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

var mockMoviesService = test_helpers.NewMockMovieService()

const title = "Title"

func setMovieCtx(movie *models.Movie) context.Context {
	return middlewares.SetMovieCtx(context.Background(), movie)
}

func TestMoviesController_HandleGetAllMovies(t *testing.T) {
	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodGet, "/movies", nil)

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetAllMovies, nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response []movieDTO

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	movies, err := mockMoviesService.GetAll(context.Background())

	if err != nil {
		t.Errorf("Mock Movie Servie Get all should not return err %v", err)
	}

	if len(response) != len(movies) {
		t.Errorf("Service Returned and the Response returned should match %v %v", len(response), len(movies))
	}
}

func TestMoviesController_HandleGetMovie(t *testing.T) {
	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodGet, "/movies/1", nil)
	movie := &models.Movie{ID: uuid.New(), Title: title}
	ctx := setMovieCtx(movie)

	rr := test_helpers.ExecuteRequest(req, controller.HandleGetMovie, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response movieDTO

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.ID != movie.ID {
		t.Errorf("ID does not match: %v %v", response.ID, movie.ID)
	}

	if response.Title != movie.Title {
		t.Errorf("Title does not match: %v %v", response.Title, title)
	}
}

func TestMoviesController_HandleCreateMovie(t *testing.T) {
	createParam := CreateMovieRequest{
		Title:       title,
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

	var response movieDTO

	if err = json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.Title != createParam.Title {
		t.Errorf("Title does not match: %v %v", response.Title, title)
	}
}

func TestMoviesController_HandleUpdateMovie(t *testing.T) {
	updateParam := UpdateMovieRequest{
		Title:       title,
		Director:    "Director",
		TicketPrice: 2,
	}

	jsonData, err := json.Marshal(updateParam)

	if err != nil {
		t.Errorf("Error encoding JSON: %v", err)
	}

	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodPut, "/movies/1", bytes.NewBuffer(jsonData))
	movie := &models.Movie{ID: uuid.New(), Title: updateParam.Title}
	ctx := setMovieCtx(movie)

	rr := test_helpers.ExecuteRequest(req, controller.HandleUpdateMovie, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response movieDTO

	if err = json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.ID != movie.ID {
		t.Errorf("ID does not match: %v %v", response.Title, title)
	}

	if response.Title != movie.Title {
		t.Errorf("Title does not match: %v %v", response.Title, title)
	}
}

func TestMoviesController_HandleDeleteMovie(t *testing.T) {
	controller := NewMoviesController(mockMoviesService)
	req := test_helpers.NewRequest(t, http.MethodDelete, "/movies/1", nil)
	movie := &models.Movie{ID: uuid.New(), Title: title}
	ctx := setMovieCtx(movie)

	rr := test_helpers.ExecuteRequest(req, controller.HandleDeleteMovie, ctx)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}
