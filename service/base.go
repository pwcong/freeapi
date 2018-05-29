package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pwcong/freeapi/config"
)

type BaseService struct {
	Conf *config.Config
	DB   *gorm.DB
}

func ConvertPageParameter(pageNo int, pageSize int) (int, int) {

	offset := (pageNo - 1) * pageSize

	if offset <= 0 {
		offset = -1
	}

	limit := pageSize
	if limit < 0 {
		limit = -1
	}

	return offset, limit
}
