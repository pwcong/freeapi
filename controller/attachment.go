package controller

import (
	"github.com/labstack/echo"
	"github.com/pwcong/freeapi/service"
)

type AttachmentController struct {
	Base *BaseController
}

func (ctx *AttachmentController) Upload(c echo.Context) error {

	service := service.AttachmentService{Base: ctx.Base.Service}

	file, err := c.FormFile("file")
	if file == nil || err != nil {
		return BaseResponse(c, false, STATUS_ERROR, err.Error(), struct{}{})
	}

	res, err := service.SaveAttachment(file)

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "upload successfully", struct {
		URL string `json:"url"`
	}{
		URL: "/public/attachment/" + res.Year + "/" + res.Month + "/" + res.Date + "/" + res.Symbol + "." + res.ExtName,
	})

}
