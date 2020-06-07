package controller

import (
	"github.com/labstack/echo/v4"
)

type QRCodeController struct {
	Base *BaseController
}

func (ctx *QRCodeController) Encode(c echo.Context) error {
	return nil
}

func (ctx *QRCodeController) Decode(c echo.Context) error {
	return nil
}
