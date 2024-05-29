package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MockMovieRepository struct {
	GetAllFunc     func() ([]*models.Movie, error)
	GetByIDFunc    func(id uuid.UUID) (*models.Movie, error)
	SaveFunc       func(param models.CreateMovieParam) (*models.Movie, error)
	UpdateByIDFunc func(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error)
	DeleteByIDFunc func(id uuid.UUID) error
}

func (m *MockMovieRepository) GetAll(_ context.Context) ([]*models.Movie, error) {
	return m.GetAllFunc()
}

func (m *MockMovieRepository) GetByID(_ context.Context, id uuid.UUID) (*models.Movie, error) {
	return m.GetByIDFunc(id)
}

func (m *MockMovieRepository) Save(_ context.Context, param models.CreateMovieParam) (*models.Movie, error) {
	return m.SaveFunc(param)
}

func (m *MockMovieRepository) UpdateByID(
	_ context.Context,
	id uuid.UUID,
	param models.UpdateMovieParam) (*models.Movie, error) {
	return m.UpdateByIDFunc(id, param)
}

func (m *MockMovieRepository) DeleteByID(_ context.Context, id uuid.UUID) error {
	return m.DeleteByIDFunc(id)
}

func TestGetAll(t *testing.T) {
	expectedMovies := []*models.Movie{
		{ID: uuid.New(), Title: "Movie 1"},
		{ID: uuid.New(), Title: "Movie 2"},
	}

	mockRepo := &MockMovieRepository{
		GetAllFunc: func() ([]*models.Movie, error) {
			return expectedMovies, nil
		},
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

	mockRepo := &MockMovieRepository{
		GetByIDFunc: func(movieID uuid.UUID) (*models.Movie, error) {
			if movieID == id {
				return expectedMovie, nil
			}
			return nil, errors.New("movie not found")
		},
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

	mockRepo := &MockMovieRepository{
		SaveFunc: func(param models.CreateMovieParam) (*models.Movie, error) {
			return expectedMovie, nil
		},
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

	mockRepo := &MockMovieRepository{
		UpdateByIDFunc: func(movieID uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
			if movieID == id {
				return expectedMovie, nil
			}
			return nil, errors.New("movie not found")
		},
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

	mockRepo := &MockMovieRepository{
		DeleteByIDFunc: func(movieID uuid.UUID) error {
			if movieID == id {
				return nil
			}
			return errors.New("movie not found")
		},
	}

	service := NewMoviesServiceImpl(mockRepo)

	err := service.Delete(context.Background(), id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
