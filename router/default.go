package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/freeapi/config"
	"github.com/pwcong/freeapi/controller"
	"github.com/pwcong/freeapi/service"
)

func Init(e *echo.Echo, conf *config.Config, db *gorm.DB) {

	e.Static("/", "view/dist")
	e.Static("/public", "public")
	e.Static("/doc", "doc")

	baseService := &service.BaseService{Conf: conf, DB: db}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}

	attachmentController := &controller.AttachmentController{Base: baseController}
	areaController := &controller.AreaController{Base: baseController}

	apiGroup := e.Group("/api")
	apiGroup.POST("/attachment/upload", attachmentController.Upload)
	apiGroup.GET("/area/:id", areaController.GetAreaByID)
	apiGroup.GET("/areas", areaController.GetAreas)

}
