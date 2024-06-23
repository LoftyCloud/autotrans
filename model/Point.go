package model

import (
	"autotrans/utils/errmsg"

	"gorm.io/gorm"
)

// 装卸点 表
// PointID	INT	装卸点ID（主键）
// PointName	VARCHAR	装卸点名
type Point struct {
	gorm.Model
	PointName string `gorm:"type:varchar(20);not null" json:"pointname"`
}

// 方法实现
// 新增装卸点，结构体为引入变量，传入方法时使用指针
func CreatePoi(data *Point) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询装卸点是否存在
func CheckPoi(PointName string) int {
	var poi Point
	db.Select("id").Where("point_name = ?", PointName).First(&poi)
	if poi.ID > 0 {
		return errmsg.ERROR_POINT_EXIST
	}
	return errmsg.SUCCESS
}

// 列表查询,一般涉及分页
func GetPoint(pageSize int, pageNum int) []Point {
	var poi []Point
	var err error

	if pageSize == -1 && pageNum == -1 {
		err = db.Limit(-1).Offset(-1).Find(&poi).Error
	} else {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&poi).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return poi
}

// 编辑装卸点信息
func EditPoi(id int, data *Point) int {
	var poi Point
	var maps = make(map[string]interface{})
	maps["point_name"] = data.PointName

	err = db.Model(&poi).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CheckPoiByID(id int) int {
	var poi Point
	db.Where("id = ?", id).First(&poi)
	if poi.ID == 0 {
		return errmsg.ERROR_POINT_NOT_EXIST
	}
	return errmsg.SUCCESS
}

// 删除装卸点
func DelPoi(id int) int {
	var poi Point
	// 软删除
	// 检查id是否存在
	if CheckPoiByID(id) != errmsg.SUCCESS {
		return errmsg.ERROR_POINT_NOT_EXIST
	}

	// 删除point
	err := db.Where("id = ?", id).Delete(&poi).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
