//+build wireinject

package main

import (
	"myapp/controller"
	"myapp/driver"
	"myapp/infrastructure/persistence"
	"myapp/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApplication() *echo.Echo {
	wire.Build(
		driver.InitDB,
		driver.NewValidator,
		persistence.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
		NewApplication,
	)
	return nil
}
