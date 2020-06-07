package controller

import (
	"errors"

	"freeapi/utils"

	"github.com/labstack/echo/v4"
)

type CodeController struct {
	Base *BaseController
}

type CodeForm struct {
	Method string `json:"method" form:"method"`
	Source string `json:"source" form:"source"`
}

func (ctx *CodeController) ConvertCoding(c echo.Context) error {

	form := new(CodeForm)

	if err := c.Bind(form); err != nil {
		return err
	}

	var res string
	var err error

	if form.Source == "" || form.Method == "" {
		return errors.New("method and source can not be empty")
	}

	if form.Method == "unicode_encode" {
		res, err = utils.UnicodeEncode(form.Source)
	} else if form.Method == "unicode_decode" {
		res, err = utils.UnicodeDecode(form.Source)
	} else if form.Method == "url_encode" {
		res, err = utils.URLEncode(form.Source)
	} else if form.Method == "url_decode" {
		res, err = utils.URLDecode(form.Source)
	} else {
		return errors.New("unknown convert method")
	}

	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "convert successfully", struct {
		Result string `json:"result"`
	}{
		Result: res,
	})

}
