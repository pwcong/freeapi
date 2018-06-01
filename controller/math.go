package controller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pwcong/freeapi/utils"
)

type MathController struct {
	Base *BaseController
}

type MathForm struct {
	Method string `json:"method" form:"method"`
	Source string `json:"source" form:"source"`
}

func (ctx *MathController) ConvertHexadecimal(c echo.Context) error {

	form := new(MathForm)

	if err := c.Bind(form); err != nil {
		return err
	}

	var res string
	var err error

	if form.Source == "" || form.Method == "" {
		return errors.New("method and source can not be empty")
	}

	if form.Method == "b2o" {
		res, err = utils.BToO(form.Source)
	} else if form.Method == "o2b" {
		res, err = utils.OToB(form.Source)
	} else if form.Method == "b2d" {
		res, err = utils.BToD(form.Source)
	} else if form.Method == "d2b" {
		res, err = utils.DToB(form.Source)
	} else if form.Method == "b2x" {
		res, err = utils.BToX(form.Source)
	} else if form.Method == "x2b" {
		res, err = utils.XToB(form.Source)
	} else if form.Method == "o2d" {
		res, err = utils.OToD(form.Source)
	} else if form.Method == "d2o" {
		res, err = utils.DToO(form.Source)
	} else if form.Method == "o2x" {
		res, err = utils.OToX(form.Source)
	} else if form.Method == "x2o" {
		res, err = utils.XToO(form.Source)
	} else if form.Method == "d2x" {
		res, err = utils.DToX(form.Source)
	} else if form.Method == "x2d" {
		res, err = utils.XToD(form.Source)
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
