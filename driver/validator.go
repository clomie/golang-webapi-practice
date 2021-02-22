package driver

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewValidator() echo.Validator {
	return &customValidator{validator: validator.New()}
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
