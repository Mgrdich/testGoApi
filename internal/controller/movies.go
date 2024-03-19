package controller

import (
	"errors"
	"net/http"
	"time"

	".com/internal/db"
	".com/internal/models"
	".com/internal/server"
	".com/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
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
	moviesStore db.MoviesStore
}

func NewMoviesController(store db.MoviesStore) *MoviesController {
	return &MoviesController{
		moviesStore: store,
	}
}

func (mC *MoviesController) HandleGetMovie(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		render.Render(w, r, server.ErrorBadRequest)
		return
	}

	movie, err := mC.moviesStore.GetByID(id)
	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, server.ErrorNotFound)
			return
		}
		render.Render(w, r, server.ErrorInternalServerError)
	}

	mr := newMovieDTO(movie)
	render.Render(w, r, mr)
}

func (mC *MoviesController) HandleGetAllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := mC.moviesStore.GetAll()
	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, server.ErrorNotFound)
			return
		}
		render.Render(w, r, server.ErrorInternalServerError)
	}

	var moviesDTO []render.Renderer
	for _, movie := range movies {
		moviesDTO = append(moviesDTO, newMovieDTO(&movie))
	}

	err = render.RenderList(w, r, moviesDTO)
	if err != nil {
		render.Render(w, r, server.ErrorConflict(err))
		return
	}
}

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

func (mC *MoviesController) HandleCreateMovie(w http.ResponseWriter, r *http.Request) {
	data := &CreateMovieRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, server.ErrorBadRequest)
		return
	}

	movie, err := mC.moviesStore.Create(models.CreateMovieParam{
		ID:          uuid.New(),
		Title:       data.Title,
		Director:    data.Director,
		ReleaseDate: time.Now().UTC(),
		TicketPrice: data.TicketPrice,
	})

	if err != nil {
		var dupKetErr *util.DuplicateKeyError
		if errors.As(err, &dupKetErr) {
			render.Render(w, r, server.ErrorConflict(err))
			return
		}

		render.Render(w, r, server.ErrorInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, newMovieDTO(movie))
}

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

func (mC *MoviesController) HandleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		render.Render(w, r, server.ErrorBadRequest)
		return
	}

	data := &UpdateMovieRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, server.ErrorBadRequest)
		return
	}

	movie, err := mC.moviesStore.Update(id, models.UpdateMovieParam{
		Title:       data.Title,
		Director:    data.Director,
		ReleaseDate: time.Now().UTC(),
		TicketPrice: data.TicketPrice,
	})

	if err != nil {
		var rnfError *util.RecordNotFoundError
		if errors.As(err, &rnfError) {
			// TODO fix it later This is wrong
			render.Render(w, r, server.ErrorNotFound)
			return
		}
		render.Render(w, r, server.ErrorInternalServerError)
		return
	}

	render.Render(w, r, newMovieDTO(movie))
}

func (mC *MoviesController) HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		render.Render(w, r, server.ErrorBadRequest)
		return
	}

	err = mC.moviesStore.Delete(id)
	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			render.Render(w, r, server.ErrorNotFound)
			return
		}
		render.Render(w, r, server.ErrorInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(nil)
}
