package controller

import (
	"errors"
	"log"
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
	PersonService services.PersonService
}

func NewPersonController(personService services.PersonService) *PersonController {
	return &PersonController{
		PersonService: personService,
	}
}

// HandleGetPerson retrieves a person by ID
// @Summary Get a person by ID
// @Description Retrieves a person by their ID
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Success 200 {object} personDTO
// @Failure 400 {object} server.HTTPError
// @Failure 404 {object} server.HTTPError
// @Router /api/v1/person/{id} [get]
func (pC *PersonController) HandleGetPerson(w http.ResponseWriter, r *http.Request) {
	person, ok := middlewares.GetPersonCtx(r.Context())

	if !ok {
		log.Println(util.NewContextCouldNotBeFetchedError())

		_ = render.Render(w, r, server.ErrorNotFound)

		return
	}

	mr := newPersonDTO(person)
	_ = render.Render(w, r, mr)
}

// HandleGetAllPerson retrieves all persons
// @Summary Get all people
// @Description Retrieves all people
// @Tags person
// @Accept  json
// @Produce  json
// @Success 200 {array} personDTO
// @Failure 404 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/person [get]
func (pC *PersonController) HandleGetAllPerson(w http.ResponseWriter, r *http.Request) {
	people, err := pC.PersonService.GetAll(r.Context())
	if err != nil {
		log.Println(err)

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
		log.Println(err)

		_ = render.Render(w, r, server.ErrorConflict(err))

		return
	}
}

// CreatePersonRequest represents the request payload for creating a person
// @Description Create Person Request
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

// HandleCreatePerson creates a new person
// @Summary Create a new person
// @Description Creates a new person with the provided data
// @Tags person
// @Accept json
// @Produce json
// @Param data body CreatePersonRequest true "Person data"
// @Success 201 {object} personDTO
// @Failure 400 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/person [post]
func (pC *PersonController) HandleCreatePerson(w http.ResponseWriter, r *http.Request) {
	data := &CreatePersonRequest{}
	if err := render.Bind(r, data); err != nil {
		log.Println(err)

		_ = render.Render(w, r, server.ErrorBadRequest)

		return
	}

	person, err := pC.PersonService.Create(r.Context(), models.CreatePerson{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	})

	if err != nil {
		log.Println(err)

		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, newPersonDTO(person))
}
