package controller

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
)

type mockMovieService struct{}

func newMockMovieService() *mockMovieService {
	return &mockMovieService{}
}

func (*mockMovieService) GetAll() ([]*models.Movie, error) {
	return nil, nil
}

func (*mockMovieService) Get(id uuid.UUID) (*models.Movie, error) {
	return nil, nil
}

func (*mockMovieService) Create(param models.CreateMovieParam) (*models.Movie, error) {
	return nil, nil
}

func (*mockMovieService) Delete(id uuid.UUID) error {
	return nil
}

func (*mockMovieService) Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	return nil, nil
}

func setMovieCtx() context.Context {
	return middlewares.SetMovieCtx(context.Background(), &models.Movie{})
}

func TestMoviesController_HandleGetAllMovies(t *testing.T) {
	controller := NewMoviesController(newMockMovieService())
	req := NewRequest(t, http.MethodGet, "/movies", nil)

	rr := ExecuteRequest(req, controller.HandleGetAllMovies, nil)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleGetMovie(t *testing.T) {
	controller := NewMoviesController(newMockMovieService())
	req := NewRequest(t, http.MethodGet, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleGetMovie, ctx)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleCreateMovie(t *testing.T) {
	controller := NewMoviesController(newMockMovieService())
	req := NewRequest(t, http.MethodPost, "/movies", nil)
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleCreateMovie, ctx)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleUpdateMovie(t *testing.T) {
	controller := NewMoviesController(newMockMovieService())
	req := NewRequest(t, http.MethodPut, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleUpdateMovie, ctx)

	CheckStatusOK(t, rr)
}

func TestMoviesController_HandleDeleteMovie(t *testing.T) {
	controller := NewMoviesController(newMockMovieService())
	req := NewRequest(t, http.MethodDelete, "/movies/1", nil)
	ctx := setMovieCtx()

	rr := ExecuteRequest(req, controller.HandleDeleteMovie, ctx)

	CheckStatusOK(t, rr)
}
