package services

import (
	"sync"
	"time"

	".com/internal/models"
	".com/internal/util"
	"github.com/google/uuid"
)

type MemoryMoviesStore struct {
	movies map[uuid.UUID]models.Movie
	mu     sync.RWMutex
}

func NewMemoryMoviesStore() *MemoryMoviesStore {
	return &MemoryMoviesStore{
		movies: make(map[uuid.UUID]models.Movie),
		mu:     sync.RWMutex{},
	}
}

func (s *MemoryMoviesStore) GetAll() ([]models.Movie, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var movies []models.Movie
	for _, m := range s.movies {
		movies = append(movies, m)
	}

	return movies, nil
}

func (s *MemoryMoviesStore) GetByID(id uuid.UUID) (*models.Movie, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	m, ok := s.movies[id]
	if !ok {
		return nil, &util.RecordNotFoundError{}
	}

	return &m, nil
}

func (s *MemoryMoviesStore) Create(param models.CreateMovieParam) (*models.Movie, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.movies[param.ID]; ok {
		return nil, &util.DuplicateKeyError{
			ID: param.ID,
		}
	}

	movie := models.Movie{
		ID:          param.ID,
		Title:       param.Title,
		Director:    param.Director,
		ReleaseDate: param.ReleaseDate,
		TicketPrice: param.TicketPrice,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	s.movies[movie.ID] = movie

	return &movie, nil
}

func (s *MemoryMoviesStore) Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	m, ok := s.movies[id]
	if !ok {
		return nil, &util.RecordNotFoundError{}
	}

	m.Title = param.Title
	m.Director = param.Director
	m.ReleaseDate = param.ReleaseDate
	m.TicketPrice = param.TicketPrice
	m.UpdatedAt = time.Now().UTC()

	s.movies[id] = m

	return &m, nil
}

func (s *MemoryMoviesStore) Delete(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.movies[id]; !ok {
		return &util.RecordNotFoundError{}
	}

	delete(s.movies, id)

	return nil
}
