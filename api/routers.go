package api

import (
	"net/http"
	ControllersUser "serotonin/api/v1/controllers/users"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userController *ControllersUser.UsersController) {
	user := e.Group("v1/users")
	user.POST("/login", userController.Login)
	user.POST("/", userController.Register)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})
}
