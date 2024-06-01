package services

import (
	"context"
	"errors"
	"testGoApi/internal/test_helpers"
	"testing"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

func TestGetAll(t *testing.T) {
	expectedMovies := []*models.Movie{
		{ID: uuid.New(), Title: "Movie 1"},
		{ID: uuid.New(), Title: "Movie 2"},
	}

	GetAllFunc := func() ([]*models.Movie, error) {
		return expectedMovies, nil
	}

	mockRepo := &test_helpers.MockMoviesRepository{
		GetAllFunc: GetAllFunc,
	}

	service := NewMoviesServiceImpl(mockRepo)

	movies, err := service.GetAll(context.Background())

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(movies) != len(expectedMovies) {
		t.Fatalf("expected %d movies, got %d", len(expectedMovies), len(movies))
	}
}

func TestGet(t *testing.T) {
	id := uuid.New()
	expectedMovie := &models.Movie{ID: id, Title: "Movie"}
	GetByIDFunc := func(movieID uuid.UUID) (*models.Movie, error) {
		if movieID == id {
			return expectedMovie, nil
		}
		return nil, errors.New("movie not found")
	}

	mockRepo := &test_helpers.MockMoviesRepository{
		GetByIDFunc: GetByIDFunc,
	}

	service := NewMoviesServiceImpl(mockRepo)

	movie, err := service.Get(context.Background(), id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if movie != expectedMovie {
		t.Fatalf("expected movie %v, got %v", expectedMovie, movie)
	}
}

func TestCreate(t *testing.T) {
	createParam := models.CreateMovieParam{Title: "New Movie"}
	expectedMovie := &models.Movie{ID: uuid.New(), Title: "New Movie"}

	SaveFunc := func(param models.CreateMovieParam) (*models.Movie, error) {
		return expectedMovie, nil
	}

	mockRepo := &test_helpers.MockMoviesRepository{
		SaveFunc: SaveFunc,
	}

	service := NewMoviesServiceImpl(mockRepo)

	movie, err := service.Create(context.Background(), createParam)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if movie != expectedMovie {
		t.Fatalf("expected movie %v, got %v", expectedMovie, movie)
	}
}

func TestUpdate(t *testing.T) {
	id := uuid.New()
	updateParam := models.UpdateMovieParam{Title: "Updated Movie"}
	expectedMovie := &models.Movie{ID: id, Title: "Updated Movie"}

	UpdateByIDFunc := func(movieID uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
		if movieID == id {
			return expectedMovie, nil
		}
		return nil, errors.New("movie not found")
	}

	mockRepo := &test_helpers.MockMoviesRepository{
		UpdateByIDFunc: UpdateByIDFunc,
	}

	service := NewMoviesServiceImpl(mockRepo)

	movie, err := service.Update(context.Background(), id, updateParam)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if movie != expectedMovie {
		t.Fatalf("expected movie %v, got %v", expectedMovie, movie)
	}
}

func TestDelete(t *testing.T) {
	id := uuid.New()

	DeleteByIDFunc := func(movieID uuid.UUID) error {
		if movieID == id {
			return nil
		}
		return errors.New("movie not found")
	}

	mockRepo := &test_helpers.MockMoviesRepository{
		DeleteByIDFunc: DeleteByIDFunc,
	}

	service := NewMoviesServiceImpl(mockRepo)

	err := service.Delete(context.Background(), id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
