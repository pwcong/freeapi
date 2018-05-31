package controller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pwcong/freeapi/service"
)

type CodeController struct {
	Base *BaseController
}

type CodeForm struct {
	Method string `form:"method" query:"method"`
	Source string `form:"source" query:"source"`
}

func (ctx *CodeController) Convert(c echo.Context) error {

	service := service.CodeService{Base: ctx.Base.Service}

	form := new(CodeForm)

	if err := c.Bind(form); err != nil {
		return err
	}

	var res string
	var err error

	if form.Source == "" {
		return errors.New("source can not be empty")
	}

	if form.Method == "ascii2unicode" {
		res, err = service.ASCII2Unicode(form.Source)
	} else if form.Method == "unicode2ascii" {
		res, err = service.Unicode2ASCII(form.Source)
	} else if form.Method == "unicode2chinese" {
		res, err = service.Unicode2Chinese(form.Source)
	} else if form.Method == "chinese2unicode" {
		res, err = service.Chinese2Unicode(form.Source)
	} else if form.Method == "uft8tochinese" {
		res, err = service.UTF8ToChinese(form.Source)
	} else if form.Method == "chinese2utf8" {
		res, err = service.ChineseToUTF8(form.Source)
	} else if form.Method == "urlencode" {
		res, err = service.URLEncode(form.Source)
	} else if form.Method == "urldecode" {
		res, err = service.URLDecode(form.Source)
	}

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "convert successfully", res)

}
