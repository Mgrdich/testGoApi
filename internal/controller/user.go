package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"testGoApi/internal/middlewares"
	"testGoApi/internal/models"
	"testGoApi/internal/server"
	"testGoApi/internal/services"
	"testGoApi/internal/util"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

type tokenDTO struct {
	Token string `json:"token"`
}

// loginUserRequest represents the request payload for creating a person
// @Description Login user request
type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (loginUserRequest *loginUserRequest) Bind(r *http.Request) error {
	if len(loginUserRequest.Username) == 0 || len(loginUserRequest.Password) == 0 {
		return errors.New("missing required Fields")
	}

	return nil
}

func (hr *tokenDTO) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

// HandleLoginUser logs the user into the system
// @Summary logs the user into the system
// @Description logs the user into the system and creates a token
// @Tags user
// @Accept json
// @Produce json
// @Param data body loginUserRequest true "user login Data"
// @Success 200 {object} tokenDTO
// @Failure 400 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/user/login [post]
func (uC *UserController) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	data := &loginUserRequest{}
	if err := render.Bind(r, data); err != nil {
		log.Println(err)

		_ = render.Render(w, r, server.ErrorBadRequest)

		return
	}

	token, err := uC.UserService.Login(r.Context(), models.LoginUser{
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		log.Print(err)

		var rnfErr *util.RecordNotFoundError

		if errors.As(err, &rnfErr) {
			_ = render.Render(w, r, server.ErrorNotFound)
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	_ = render.Render(w, r, &tokenDTO{
		Token: token,
	})
}

// registerUserRequest represents the request payload for creating a person
// @Description register user request
type registerUserRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

func (registerUserRequest *registerUserRequest) Bind(r *http.Request) error {
	if len(registerUserRequest.Username) == 0 ||
		len(registerUserRequest.Password) == 0 ||
		len(registerUserRequest.RepeatPassword) == 0 {
		return errors.New("missing required Fields")
	}

	if len(registerUserRequest.Username) < 3 || len(registerUserRequest.Username) > 30 {
		return errors.New("username should be at least 4 letters and 30 letters most")
	}

	if registerUserRequest.Password != registerUserRequest.RepeatPassword {
		return errors.New("passwords do not match")
	}

	isPasswordValid, _, _, _, _ := util.ValidatePassword(registerUserRequest.Password)

	if !isPasswordValid {
		// TODO fix the whole structure
		return errors.New("password is not valid")
	}

	return nil
}

// HandleRegisterUser registers a new user
// @Summary Create a new user
// @Description Creates a new user with the provided data
// @Tags user
// @Accept json
// @Produce json
// @Param data body registerUserRequest true "user register data"
// @Success 201 {object} OKDto
// @Failure 400 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/user/register [post]
func (uC *UserController) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	data := &registerUserRequest{}

	if err := render.Bind(r, data); err != nil {
		log.Println(err)

		_ = render.Render(w, r, server.ErrorBadRequest)

		return
	}

	_, err := uC.UserService.Create(r.Context(), models.CreateUser{
		Username: data.Username,
		Password: data.Password,
		Role:     models.BasicRole,
	})

	if err != nil {
		log.Println(err)

		_ = render.Render(w, r, server.ErrorInternalServerError)

		return
	}

	render.Status(r, http.StatusCreated)

	_ = render.Render(w, r, NewOkDto())
}

type userDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}

func (hr *userDto) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

// HandleUserMe get the information of the current user
// @Summary Create a new person
// @Description Creates a new person with the provided data
// @Tags user
// @Accept json
// @Produce json
// @Param data body CreatePersonRequest true "Person data"
// @Success 200 {object} userDto
// @Failure 400 {object} server.HTTPError
// @Failure 500 {object} server.HTTPError
// @Router /api/v1/user/me [post]
func (uC *UserController) HandleUserMe(w http.ResponseWriter, r *http.Request) {
	user, ok := middlewares.GetTokenizedUserCtx(r.Context())

	if !ok {
		log.Println(util.NewContextCouldNotBeFetchedError())

		_ = render.Render(w, r, server.ErrorBadRequest)

		return
	}

	_ = render.Render(w, r, &userDto{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
}
