package main

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pwcong/freeapi/config"
	"github.com/pwcong/freeapi/controller"
	"github.com/pwcong/freeapi/db"
	"github.com/pwcong/freeapi/middleware"
	"github.com/pwcong/freeapi/model"
	"github.com/pwcong/freeapi/router"
	"github.com/pwcong/freeapi/utils"
)

func initMiddlewares(e *echo.Echo, conf *config.Config) {
	middleware.Init(e, conf)
}

func initRoutes(e *echo.Echo, conf *config.Config, db *gorm.DB) {
	router.Init(e, conf, db)
}

func initDB(db *gorm.DB) {
	db.AutoMigrate(&model.Attachment{}, &model.Area{})
}

func main() {

	rootDir := utils.GetRootDir()

	// 初始化配置
	conf, err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	dbConfig, ok := conf.Databases["sqlite3"]
	if !ok {
		log.Fatal("Can not load configuration of SQLite3")
	}

	orm := db.ORM{Name: "sqlite3"}

	orm.Open(path.Join(rootDir, dbConfig.DBPath))

	defer orm.Close()

	initDB(orm.DB)

	e := echo.New()

	// 全局错误处理
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(http.StatusOK, controller.BaseResponseJSON{
			Success: false,
			Code:    controller.STATUS_ERROR,
			Message: "url=" + c.Request().RequestURI + ", " + err.Error(),
		})
	}

	// 初始化中间件
	initMiddlewares(e, &conf)
	// 初始化路由
	initRoutes(e, &conf, orm.DB)

	// 运行服务
	e.Logger.Fatal(e.Start(conf.Server.Host + ":" + strconv.Itoa(conf.Server.Port)))

}

func init() {

}
