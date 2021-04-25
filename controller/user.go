package controller

import (
	"github.com/labstack/echo/v4"
	"golang-echo-postgresql-rest-api-example/model"
	"golang-echo-postgresql-rest-api-example/repository"
	"golang-echo-postgresql-rest-api-example/util"
	"net/http"
	"strconv"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{userRepository: userRepository}
}

// SaveUser godoc
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param user body model.UserInput true "New User"
// @Success 200 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /signup [post]
func (userController *UserController) SaveUser(c echo.Context) error {
	payload := new(model.UserInput)
	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user := &model.User{UserInput: payload}

	createdUser, err := userController.userRepository.SaveUser(user)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusCreated, createdUser)
}

// GetAllUser godoc
// @Summary Get all users
// @Description Get all user items
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(xml, json)
// @Param page query int false "page" minimum(1)
// @Param limit query int false "size" minimum(1)
// @Success 200 {array} model.User
// @Failure 500 {object} handler.APIError
// @Router /users [get]
func (userController *UserController) GetAllUser(c echo.Context) error {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64)

	pagedUser, _ := userController.userRepository.GetAllUser(page, limit)
	return util.Negotiate(c, http.StatusOK, pagedUser)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [get]
func (userController *UserController) GetUser(c echo.Context) error {
	id := c.Param("id")

	user := c.Get("user")

	user, err := userController.userRepository.GetUser(id)
	if err != nil {
		return err
	}

	return util.Negotiate(c, http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Param user body model.UserInput true "User Info"
// @Success 200 {object} model.User
// @Failure 400 {object} handler.APIError
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [put]
func (userController *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	payload := new(model.UserInput)

	if err := util.BindAndValidate(c, payload); err != nil {
		return err
	}

	user, err := userController.userRepository.UpdateUser(id, &model.User{UserInput: payload})
	if err != nil {
		return err
	}
	return util.Negotiate(c, http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a new user item
// @Tags users
// @Accept json,xml
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 204 {object} model.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{id} [delete]
func (userController *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := userController.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
