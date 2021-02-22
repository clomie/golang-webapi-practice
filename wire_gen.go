// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/labstack/echo/v4"
	"myapp/controller"
	"myapp/driver"
	"myapp/infrastructure/persistence"
	"myapp/service"
)

// Injectors from wire.go:

func initApplication() *echo.Echo {
	validator := driver.NewValidator()
	db := driver.InitDB()
	userRepository := persistence.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	echoEcho := NewApplication(validator, userController)
	return echoEcho
}
