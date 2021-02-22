package main

import (
	"myapp/controller"
	"myapp/infrastructure/persistence"
	"myapp/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func main() {
	runApp()
}

func runApp() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())

	e.Validator = &customValidator{validator: validator.New()}

	db := initDB()
	initUsers(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func initUsers(e *echo.Echo, db *gorm.DB) {
	repository := persistence.NewUserRepository(db)
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)
	e.GET("/users", controller.ListUsers)
	e.POST("/users", controller.PostUsers)
}
