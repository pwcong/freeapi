package controller

import (
	"net/http"

	"freeapi/config"
	"freeapi/service"

	"github.com/labstack/echo/v4"
)

const (
	STATUS_OK    = 20000
	STATUS_ERROR = 40000
)

type BaseController struct {
	Conf    *config.Config
	Service *service.BaseService
}

type BaseResponseJSON struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func BaseResponse(c echo.Context, success bool, code int, message string, payload interface{}) error {

	return c.JSON(http.StatusOK, BaseResponseJSON{
		Success: success,
		Code:    code,
		Message: message,
		Payload: payload,
	})

}

func ResolvePageParameter(pageNo int, pageSize int) (int, int) {

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 1 {
		pageSize = -1
	}

	return pageNo, pageSize
}
