package controller

import (
	"net/http"
	"refactoring/api/request"
	"refactoring/api/response"
	"refactoring/exceptions"
	"refactoring/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserController struct {
	userRepository repository.Repository
}

func NewUserController(repository repository.Repository) *UserController {
	return &UserController{repository}
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	request := request.CreateUserRequest{}

	err := render.Bind(r, &request)
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	userId := controller.userRepository.CreateUser(request)

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": userId,
	})
}

func (controller *UserController) DeleteUser(w http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")

	result := controller.userRepository.DeleteUser(id)

	if result == nil {
		_ = render.Render(w, request, response.ErrInvalidRequest(exceptions.UserNotFound))
		return
	}
	render.Status(request, http.StatusNoContent)
}

func (controller *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

	request := request.UpdateUserRequest{}

	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	request.Id = chi.URLParam(r, "id")

	err := controller.userRepository.UpdateUser(request)
	if err != nil {
		_ = render.Render(w, r, response.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	userList := controller.userRepository.GetAllUsers()

	id := chi.URLParam(r, "id")

	render.JSON(w, r, userList[id])
}

func (controller *UserController) SearchUsers(w http.ResponseWriter, r *http.Request) {
	userList := controller.userRepository.GetAllUsers()

	render.JSON(w, r, userList)
}
