package controllers

import (
	"net/http"
	"serotonin/business/users"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	User_service users.Services
}

func InitUserController(service users.Services) *UsersController {
	return &UsersController{
		User_service: service,
	}
}

func (controller *UsersController) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Oke": "Oke",
	})
}
