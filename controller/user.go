package controller

import (
	"myapp/domain/model"
	"myapp/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	ListUsers(c echo.Context) error
	PostUsers(c echo.Context) error
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

type listUsersRequest struct {
	Limit  int `query:"limit" validate:"min=1,max=100"`
	Offset int `query:"offset" validate:"min=0"`
}

type listUsersResult struct {
	Items  []model.User `json:"items"`
	Total  int          `json:"total"`
	Limit  int          `json:"limit"`
	Offset int          `json:"offset"`
}

func (u *userController) ListUsers(c echo.Context) error {
	req := &listUsersRequest{
		Limit:  20,
		Offset: 0,
	}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cmd := &service.ListUsersCommand{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	res, err := u.userService.ListUsers(cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := listUsersResult{
		Items:  res.Items,
		Total:  res.Total,
		Limit:  res.Limit,
		Offset: res.Offset,
	}

	return c.JSON(http.StatusOK, response)
}

type postUsersRequest struct {
	Name string `json:"name" validate:"required,min=1,max=10"`
}

type postUsersResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *userController) PostUsers(c echo.Context) error {
	req := new(postUsersRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cmd := &service.CreateUserCommand{
		Name: req.Name,
	}

	res, err := u.userService.CreateUser(cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := postUsersResponse{
		ID:        res.ID,
		Name:      res.Name,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}
