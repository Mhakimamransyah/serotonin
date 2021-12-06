package api

import (
	"net/http"
	"serotonin/api/middleware/apikey"
	"serotonin/api/v1/controllers/pilgan"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, pilgan_controller *pilgan.PilganController) {

	pilgan := e.Group("/v1/pilgan")
	pilgan.Use(apikey.ApiKey())
	pilgan.GET("", pilgan_controller.GetDataUjianPilganController)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "OK",
		})
	})
}
