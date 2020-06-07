package controller

import (
	"strconv"

	"freeapi/model"
	"freeapi/service"

	"github.com/labstack/echo/v4"
)

type AreaController struct {
	Base *BaseController
}

type AreaForm struct {
	PageSize int    `query:"page_size"`
	PageNo   int    `query:"page_no"`
	Query    string `query:"query"`
	Filter   string `query:"filter"`
}

func (ctx *AreaController) GetAreaByID(c echo.Context) error {

	service := service.AreaService{Base: ctx.Base.Service}

	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		return err
	}

	area, err := service.GetByID(uint(id))
	if err != nil {
		return err
	}

	return BaseResponse(c, true, STATUS_OK, "get area successfully", area)

}

func (ctx *AreaController) GetAreas(c echo.Context) error {

	service := service.AreaService{Base: ctx.Base.Service}

	form := new(AreaForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	pageNo, pageSize := ResolvePageParameter(form.PageNo, form.PageSize)

	var page model.Page
	var err error
	if form.Filter == "parent_id" {
		parent_id, err := strconv.Atoi(form.Query)
		if err != nil {
			return err
		}
		page, err = service.GetByParentID(uint(parent_id), pageNo, pageSize)
	} else if form.Filter == "depth" {

		depth, err := strconv.Atoi(form.Query)
		if err != nil {
			return err
		}
		page, err = service.GetByDepth(uint(depth), pageNo, pageSize)
	} else if form.Filter == "name" {
		page, err = service.QueryByName(form.Query, pageNo, pageSize)
	} else {
		page, err = service.GetByParentID(0, pageNo, pageSize)
	}

	if err != nil {
		return err
	}
	return BaseResponse(c, true, STATUS_OK, "get areas successfully", page)
}
