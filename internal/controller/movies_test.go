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
)

type mockMovieService struct{}

func (*mockMovieService) GetAll(_ context.Context) ([]*models.Movie, error) {
	return nil, nil
}

func (*mockMovieService) Get(_ context.Context, id uuid.UUID) (*models.Movie, error) {
	return &models.Movie{ID: id}, nil
}

func (*mockMovieService) Create(_ context.Context, param models.CreateMovieParam) (*models.Movie, error) {
	return &models.Movie{
		Title:       param.Title,
		Director:    param.Director,
		TicketPrice: param.TicketPrice,
	}, nil
}

func (*mockMovieService) Delete(_ context.Context, id uuid.UUID) error {
	return nil
}

func (*mockMovieService) Update(_ context.Context, id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	return &models.Movie{
		ID:          id,
		Title:       param.Title,
		Director:    param.Director,
		TicketPrice: param.TicketPrice,
	}, nil
}

func setMovieCtx() context.Context {
	return middlewares.SetMovieCtx(context.Background(), &models.Movie{})
}

func TestMoviesController_HandleGetAllMovies(t *testing.T) {
	controller := NewMoviesController(&mockMovieService{})
	req := NewRequest(t, http.MethodGet, "/movies", nil)

	rr := ExecuteRequest(req, controller.HandleGetAllMovies, nil)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleGetMovie(t *testing.T) {
	controller := NewMoviesController(&mockMovieService{})
	req := NewRequest(t, http.MethodGet, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleGetMovie, ctx)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleCreateMovie(t *testing.T) {
	createParam := CreateMovieRequest{
		Title:       "title",
		Director:    "Director",
		TicketPrice: 2,
	}

	jsonData, err := json.Marshal(createParam)

	if err != nil {
		t.Error("Error encoding JSON:", err)
	}

	controller := NewMoviesController(&mockMovieService{})

	// json content-type
	req := NewRequest(t, http.MethodPost, "/movies", bytes.NewBuffer(jsonData))
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleCreateMovie, ctx)

	CheckStatusCreated(t, rr)
}

func TestMoviesController_HandleUpdateMovie(t *testing.T) {
	updateParam := UpdateMovieRequest{
		Title:       "title",
		Director:    "Director",
		TicketPrice: 2,
	}

	jsonData, err := json.Marshal(updateParam)

	if err != nil {
		t.Error("Error encoding JSON:", err)
	}

	controller := NewMoviesController(&mockMovieService{})
	req := NewRequest(t, http.MethodPut, "/movies/1", bytes.NewBuffer(jsonData))
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleUpdateMovie, ctx)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleDeleteMovie(t *testing.T) {
	controller := NewMoviesController(&mockMovieService{})
	req := NewRequest(t, http.MethodDelete, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleDeleteMovie, ctx)

	CheckStatusOK(t, rr)
}
