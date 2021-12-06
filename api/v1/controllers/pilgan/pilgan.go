package pilgan

import (
	"encoding/json"
	"serotonin/api/common"
	"serotonin/business/pilgan"

	"github.com/labstack/echo/v4"
)

type PilganController struct {
	pilgan_service pilgan.Service
}

func InitPilganController(service pilgan.Service) *PilganController {
	return &PilganController{
		pilgan_service: service,
	}
}

func (controller *PilganController) GetDataUjianPilganController(c echo.Context) error {
	params := make(map[string]interface{})
	data, _ := json.Marshal(c.QueryParams())
	json.Unmarshal(data, &params)
	res, err := controller.pilgan_service.ReadDataPilgan(params)
	if err != nil {
		return c.JSON(common.NewBadRequestResponseWithMessage(err.Error()))
	}
	if _, exists := params["limit"]; !exists {
		params["limit"] = []interface{}{100}
	}
	if _, exists := params["skip"]; !exists {
		params["skip"] = []interface{}{0}
	}
	return c.JSON(common.NewSuccessResponseGetData(res, params, len(*res)))
}
