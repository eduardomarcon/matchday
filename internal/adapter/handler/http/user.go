package http

import (
	"api/internal/core/domain"
	"api/internal/core/port"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

type userRequest struct {
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Role     domain.UserRole `json:"role"`
}

func (ur userRequest) Bind(_ *http.Request) error {
	if ur.Name == "" {
		return errors.New("name is required")
	}
	if ur.Email == "" {
		return errors.New("email is required")
	}
	if ur.Password == "" {
		return errors.New("password is required")
	}
	if ur.Role == "" {
		return errors.New("role is required")
	}
	_, err := domain.GetUserRole(string(ur.Role))
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

type UserResponse struct {
	Name  string          `json:"name"`
	Email string          `json:"email"`
	Role  domain.UserRole `json:"role"`
}

func (ur UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserResponse(user *domain.User) *UserResponse {
	resp := &UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return resp
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	data := &userRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, handleErrBadRequest(err))
		return
	}

	user := domain.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
	createdUser, err := uh.svc.Register(r.Context(), &user)
	if err != nil {
		renderer := handleError("failed to register user", err)
		_ = render.Render(w, r, renderer)
		return
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, NewUserResponse(createdUser))
}
