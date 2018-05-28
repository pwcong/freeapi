package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pwcong/freeapi/config"
	"github.com/pwcong/freeapi/service"
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

type Page struct {
	Data      []interface{} `json:"data"`
	PageSize  int           `json:"page_size"`
	PageNo    int           `json:"page_no"`
	TotalSize int           `json:"total_size"`
	TotalNo   int           `json:"total_no"`
}

type PageForm struct {
	PageSize int `json:"page_size" form:"page_size" query:"page_size"`
	PageNo   int `json:"page_no" form:"page_no" query:"page_no"`
}
