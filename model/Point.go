package model

import (
	// "autotrans/utils/errmsg"

	"gorm.io/gorm"
)

// 装卸点 表
// PointID	INT	装卸点ID（主键）
// PointName	VARCHAR	装卸点名
type Point struct {
	gorm.Model
	PointName string `gorm:"type:varchar(20);not null" json:"name"`
}

// // 方法实现
// // 新增装卸点，结构体为引入变量，传入方法时使用指针
// func Createpoi(data *Point) int {
// 	err := db.Create(&data).Error
// 	if err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCESS
// }

// // 查询装卸点是否存在
// func Checkpoi(name string) int {
// 	var poi Point
// 	db.Select("id").Where("Name = ?", name).First(&poi)
// 	if poi.ID > 0 {
// 		return errmsg.ERROR_poi_USED
// 	}
// 	return errmsg.SUCCESS
// }

// // todo 查询装卸点下的所有篮子

// // // 查询用户列表，列表查询一般涉及分页
// // func Getpoi(pageSize int, pageNum int) []Point {
// // 	var poi []Point
// // 	var err error
// // 	if pageSize == -1 && pageNum == -1 {
// // 		err = db.Limit(-1).Offset(-1).Find(&poi).Error
// // 	} else {
// // 		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&poi).Error
// // 	}
// // 	if err != nil && err != gorm.ErrRecordNotFound {
// // 		return nil
// // 	}
// // 	return poi
// // }

// // 编辑装卸点信息（密码除外）（根据用户id）
// func Eidtpoi(id int, data *Point) int {
// 	var poi Point
// 	var maps = make(map[string]interface{})
// 	maps["name"] = data.Name

// 	err = db.Model(&poi).Where("id = ?", id).Updates(maps).Error
// 	if err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCESS
// }

// // 删除装卸点
// func Delpoi(id int) int {
// 	var poi Point
// 	// 软删除
// 	err = db.Where("id = ?", id).Delete(&poi).Error
// 	if err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCESS
// }
