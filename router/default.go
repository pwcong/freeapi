package router

import (
	"freeapi/config"
	"freeapi/controller"
	"freeapi/service"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, conf *config.Config, db *gorm.DB) {

	e.Static("/", "view/dist")
	e.Static("/public", "public")
	e.Static("/docs", "docs")

	baseService := &service.BaseService{Conf: conf, DB: db}
	baseController := &controller.BaseController{Conf: conf, Service: baseService}

	attachmentController := &controller.AttachmentController{Base: baseController}
	areaController := &controller.AreaController{Base: baseController}
	codeController := &controller.CodeController{Base: baseController}
	mathController := &controller.MathController{Base: baseController}

	apiGroup := e.Group("/api")

	// 附件接口
	apiGroup.POST("/attachment/upload", attachmentController.Upload)
	apiGroup.POST("/attachment/uploads", attachmentController.Uploads)

	// 地域接口
	apiGroup.GET("/area/:id", areaController.GetAreaByID)
	apiGroup.GET("/areas", areaController.GetAreas)

	// 编码接口
	apiGroup.POST("/code/convert/coding", codeController.ConvertCoding)

	// 数学接口
	apiGroup.POST("/math/convert/hexadecimal", mathController.ConvertHexadecimal)
}
