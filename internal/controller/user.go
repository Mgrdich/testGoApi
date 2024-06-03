package controller

import (
	"errors"
	"github.com/google/uuid"
	"net/http"

	"github.com/go-chi/render"
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

func (uC *UserController) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	data := &loginUserRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	token, err := uC.UserService.Login(r.Context(), models.LoginUser{
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		var rnfErr *util.RecordNotFoundError
		if errors.As(err, &rnfErr) {
			_ = render.Render(w, r, server.ErrorNotFound)
			return
		}

		_ = render.Render(w, r, server.ErrorInternalServerError)
	}

	_ = render.Render(w, r, &tokenDTO{
		Token: token,
	})
}

// registerUserRequest represents the request payload for creating a person
// @Description ser request
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

	if len(registerUserRequest.Username) < 3 {
		return errors.New("username should be at least 4 letters")
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

func (uC *UserController) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	data := &registerUserRequest{}

	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return
	}

	_, err := uC.UserService.Create(r.Context(), models.CreateUser{
		Username: data.Username,
		Password: data.Password,
		Role:     models.BasicRole,
	})

	if err != nil {
		_ = render.Render(w, r, server.ErrorInternalServerError)
		return
	}

	_ = render.Render(w, r, nil)
}

type userDto struct {
	ID       uuid.UUID       `json:"id"`
	Username string          `json:"username"`
	Role     models.UserRole `json:"role"`
}

func (hr *userDto) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (uC *UserController) HandleUserMe(w http.ResponseWriter, r *http.Request) {

}
