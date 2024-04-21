package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"testGoApi.com/internal/db"
	"testGoApi.com/internal/models"
	"testGoApi.com/internal/server"
	"testGoApi.com/internal/util"
)

type personDTO struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

func newPersonDTO(p *models.Person) *personDTO {
	return &personDTO{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
}

func (hr *personDTO) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PersonController struct {
	personStore db.PersonStore
}

func NewPersonController(store db.PersonStore) *PersonController {
	return &PersonController{
		personStore: store,
	}
}

func (pC *PersonController) HandleGetPerson(w http.ResponseWriter, r *http.Request) {
	personId := chi.URLParam(r, "id")
	if personId == "" {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	id, err := uuid.Parse(personId)

	if err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	person, err := pC.personStore.GetByID(id)

	if err != nil {
		_ = render.Render(w, r, server.ErrorNotFound)
		return
	}

	mr := newPersonDTO(person)
	_ = render.Render(w, r, mr)
}

func (pC *PersonController) HandleGetAllPerson(w http.ResponseWriter, r *http.Request) {
	people, err := pC.personStore.GetAll()
	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			_ = render.Render(w, r, server.ErrorNotFound)
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)
	}

	var peopleDTO []render.Renderer
	for i := range people {
		peopleDTO = append(peopleDTO, newPersonDTO(people[i]))
	}

	err = render.RenderList(w, r, peopleDTO)
	if err != nil {
		_ = render.Render(w, r, server.ErrorConflict(err))
		return
	}
}

type CreatePersonRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (pr *CreatePersonRequest) Bind(r *http.Request) error {
	if len(pr.FirstName) == 0 || len(pr.LastName) == 0 {
		return errors.New("missing required Fields")
	}

	return nil
}

func (pC *PersonController) HandleCreatePerson(w http.ResponseWriter, r *http.Request) {
	data := &CreatePersonRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	person, err := pC.personStore.Create(models.CreatePerson{
		ID:        uuid.New(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
	})

	if err != nil {
		var dupKetErr *util.DuplicateKeyError
		if errors.As(err, &dupKetErr) {
			_ = render.Render(w, r, server.ErrorConflict(err))
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, newPersonDTO(person))
}
