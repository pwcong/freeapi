package controller

import (
	"freeapi/service"

	"github.com/labstack/echo/v4"
)

type AttachmentController struct {
	Base *BaseController
}

func (ctx *AttachmentController) Upload(c echo.Context) error {

	service := service.AttachmentService{Base: ctx.Base.Service}

	file, err := c.FormFile("file")
	if file == nil || err != nil {
		return err
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

func (ctx *AttachmentController) Uploads(c echo.Context) error {

	service := service.AttachmentService{Base: ctx.Base.Service}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]

	var urls []string

	for _, file := range files {

		res, err := service.SaveAttachment(file)

		if err == nil {
			urls = append(urls, "/public/attachment/"+res.Year+"/"+res.Month+"/"+res.Date+"/"+res.Symbol+"."+res.ExtName)
		}

	}

	return BaseResponse(c, true, STATUS_OK, "upload successfully", struct {
		URLS []string `json:"urls"`
	}{
		URLS: urls,
	})

}
