package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testGoApi/internal/server"
	"testGoApi/internal/services"
	"testGoApi/internal/util"
)

type movieDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newMovieDTO(m *models.Movie) *movieDTO {
	return &movieDTO{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (hr *movieDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type MoviesController struct {
	MovieService services.MovieService
}

func NewMoviesController(store services.MovieService) *MoviesController {
	return &MoviesController{
		MovieService: store,
	}
}

// HandleGetMovie retrieves a movie by its context
// @Summary Get a movie by context
// @Description Retrieves a movie using the context set by middleware
// @Tags movie
// @Param id path string true "Movie ID"
// @Accept json
// @Produce json
// @Success 200 {object} movieDTO
// @Failure 400 {object} server.HTTPError
// @Router /api/v1/movies/{id} [get]
func (mC *MoviesController) HandleGetMovie(w http.ResponseWriter, r *http.Request) {
	movie, ok := middlewares.GetMovieCtx(r.Context())
	if !ok {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	mr := newMovieDTO(movie)
	_ = render.Render(w, r, mr)
}

// HandleGetAllMovies retrieves all movies
// @Summary Get all movies
// @Description Retrieves all movies from the database
// @Tags movie
// @Accept json
// @Produce json
// @Success 200 {array} movieDTO
// @Failure 404 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/movies [get]
func (mC *MoviesController) HandleGetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := mC.MovieService.GetAll()
	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			_ = render.Render(w, r, server.ErrorNotFound)
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)
	}

	var moviesDTO []render.Renderer
	for i := range movies {
		moviesDTO = append(moviesDTO, newMovieDTO(movies[i]))
	}

	err = render.RenderList(w, r, moviesDTO)
	if err != nil {
		_ = render.Render(w, r, server.ErrorConflict(err))
		return
	}
}

// CreateMovieRequest represents the request payload for creating a movie
// @Description Create Movie Request
type CreateMovieRequest struct {
	Title       string  `json:"title"`
	Director    string  `json:"director"`
	TicketPrice float64 `json:"ticket_price"`
}

func (mr *CreateMovieRequest) Bind(r *http.Request) error {
	if len(mr.Title) == 0 || len(mr.Director) == 0 {
		return errors.New("missing required Fields")
	}

	return nil
}

// HandleCreateMovie creates a new movie
// @Summary Create a new movie
// @Description Creates a new movie with the provided details
// @Tags movie
// @Accept json
// @Produce json
// @Param movie body CreateMovieRequest true "Create Movie Request"
// @Success 201 {object} movieDTO
// @Failure 400 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/movies [post]
func (mC *MoviesController) HandleCreateMovie(w http.ResponseWriter, r *http.Request) {
	data := &CreateMovieRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	movie, err := mC.MovieService.Create(models.CreateMovieParam{
		Title:       data.Title,
		Director:    data.Director,
		ReleaseDate: time.Now().UTC(),
		TicketPrice: data.TicketPrice,
	})

	if err != nil {
		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, newMovieDTO(movie))
}

// UpdateMovieRequest represents the request payload for updating a movie
// @Description Update Movie Request
type UpdateMovieRequest struct {
	Title       string  `json:"title"`
	Director    string  `json:"director"`
	TicketPrice float64 `json:"ticket_price"`
}

func (mr *UpdateMovieRequest) Bind(r *http.Request) error {
	if len(mr.Title) == 0 || len(mr.Director) == 0 {
		return errors.New("missing required Fields")
	}

	return nil
}

// HandleUpdateMovie updates an existing movie
// @Summary Update an existing movie
// @Description Updates an existing movie with the provided details
// @Tags movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Param movie body UpdateMovieRequest true "Update Movie Request"
// @Success 200 {object} movieDTO
// @Failure 400 {object} server.HTTPError
// @Failure 404 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/movies/{id} [put]
func (mC *MoviesController) HandleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	movie, ok := middlewares.GetMovieCtx(r.Context())
	if !ok {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	data := &UpdateMovieRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	updatedMovie, err := mC.MovieService.Update(movie.ID, models.UpdateMovieParam{
		Title:       data.Title,
		Director:    data.Director,
		ReleaseDate: time.Now().UTC(),
		TicketPrice: data.TicketPrice,
	})

	if err != nil {
		var rnfError *util.RecordNotFoundError
		if errors.As(err, &rnfError) {
			_ = render.Render(w, r, server.ErrorNotFound)
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	_ = render.Render(w, r, newMovieDTO(updatedMovie))
}

// HandleDeleteMovie deletes a movie by ID
// @Summary Delete a movie by ID
// @Description Deletes a movie by its ID
// @Tags movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200
// @Failure 400 {object} server.HTTPError
// @Failure 404 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/movies/{id} [delete]
func (mC *MoviesController) HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	movie, ok := middlewares.GetMovieCtx(r.Context())
	if !ok {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	err := mC.MovieService.Delete(movie.ID)
	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			_ = render.Render(w, r, server.ErrorNotFound)
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(nil)

	if err != nil {
		_ = render.Render(w, r, server.ErrorInternalServerError)
		return
	}
}
