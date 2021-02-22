package main

import (
	"myapp/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// resolve dependencies by wire
	e := initApplication()

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}

func NewApplication(
	validator echo.Validator,
	userController controller.UserController,
) *echo.Echo {
	e := echo.New()

	// logger
	e.Logger.SetLevel(log.DEBUG)

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// validator
	e.Validator = validator

	// routings
	e.GET("/users", userController.ListUsers)
	e.POST("/users", userController.PostUsers)
	e.GET("/users/:id", userController.GetUser)
	e.PUT("/users/:id", userController.PutUser)
	e.DELETE("/users/:id", userController.DeleteUser)

	return e
}
