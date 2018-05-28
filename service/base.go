package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pwcong/freeapi/config"
)

type BaseService struct {
	Conf *config.Config
	DB   *gorm.DB
}
