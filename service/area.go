package service

import (
	"errors"

	"freeapi/model"
)

type AreaService struct {
	Base *BaseService
}

func (ctx *AreaService) GetByID(id uint) (model.Area, error) {

	db := ctx.Base.DB

	var area model.Area

	if notFound := db.Where("id = ?", id).First(&area).RecordNotFound(); notFound {
		return model.Area{}, errors.New("area is not existed")
	}

	return area, nil
}

func (ctx *AreaService) GetByParentID(id uint, pageNo int, pageSize int) (model.Page, error) {
	db := ctx.Base.DB

	var totalSize int
	if err := db.Where("parent_id = ?", id).Find(&[]model.Area{}).Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var areas []model.Area

	offset, limit := ConvertPageParameter(pageNo, pageSize)

	if err := db.Where("parent_id = ?", id).Offset(offset).Limit(limit).Find(&areas).Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(areas),
		TotalSize:   totalSize,
		Data:        areas,
	}, nil

}

func (ctx *AreaService) GetByDepth(depth uint, pageNo int, pageSize int) (model.Page, error) {
	db := ctx.Base.DB

	var totalSize int
	if err := db.Where("depth = ?", depth).Find(&[]model.Area{}).Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var areas []model.Area
	offset, limit := ConvertPageParameter(pageNo, pageSize)
	if err := db.Where("depth = ?", depth).Offset(offset).Limit(limit).Find(&areas).Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(areas),
		TotalSize:   totalSize,
		Data:        areas,
	}, nil
}

func (ctx *AreaService) QueryByName(name string, pageNo int, pageSize int) (model.Page, error) {

	db := ctx.Base.DB
	var totalSize int

	if err := db.Where("name like ?", "%"+name+"%").Find(&[]model.Area{}).Count(&totalSize).Error; err != nil {
		return model.Page{}, err
	}

	var areas []model.Area
	offset, limit := ConvertPageParameter(pageNo, pageSize)

	if err := db.Where("name like ?", "%"+name+"%").Offset(offset).Limit(limit).Find(&areas).Error; err != nil {
		return model.Page{}, err
	}

	return model.Page{
		PageNo:      pageNo,
		PageSize:    pageSize,
		CurrentSize: len(areas),
		TotalSize:   totalSize,
		Data:        areas,
	}, nil
}
