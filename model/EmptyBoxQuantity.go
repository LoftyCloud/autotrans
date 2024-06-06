package model

import (
	// "autotrans/utils/errmsg"

	"gorm.io/gorm"
)

// 空料框数量记录表
// RecordID	INT	记录ID（主键）
// PointID	INT	装卸点ID（外键，关联装卸点表）
// DateTime	DATETIME	记录时间
// EmptyBoxes	INT	空料框数量
type EmptyBoxQuantity struct {
	Point Point `gorm:"foreignkey:PointID"`  // 外键
	gorm.Model
	PointID   int    `gorm:"type:int;not null" json:"pointID"`
	EmptyBoxes   int `gorm:"type:int;not null" json:"num"`
}

// // 方法实现
// // 新增框，结构体为引入变量，传入方法时使用指针
// func CreateEmptyBoxQuantity(data *EmptyBoxQuantity) int {
// 	err := db.Create(&data).Error
// 	if err != nil {
// 		// fmt.Println(err)
// 		return errmsg.ERROR_CATE_NOT_EXIST
// 	}
// 	return errmsg.SUCCESS
// }

// // todo 查询单个装卸点下的所有空料框
// func GetCateEbq(id int, pageSize int, pageNum int) ([]EmptyBoxQuantity, int) {
// 	var cateEbqList []EmptyBoxQuantity
// 	// var total int64
// 	if pageSize == -1 && pageNum == -1 {
// 		err = db.Preload("Point").Limit(-1).Offset(-1).Where("cid=?", id).Find(&cateEbqList).Error
// 	} else {
// 		err = db.Preload("Point").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateEbqList).Error
// 	}
// 	// db.Model(&cateEbqList).Where("cid =?", id).Count(&total)
// 	if err != nil {
// 		return nil, errmsg.ERROR_CATE_NOT_EXIST
// 	}
// 	return cateEbqList, errmsg.SUCCESS
// }

// // todo 查询单个文章
// func GetEbqInfo(id int) (EmptyBoxQuantity, int) {
// 	var Ebq EmptyBoxQuantity
// 	err = db.Where("id=?", id).Preload("Point").First(&Ebq).Error
// 	if err != nil {
// 		return Ebq, errmsg.ERROR_Ebq_NOT_EXIST
// 	}
// 	return Ebq, errmsg.SUCCESS
// }

// // 查询文章列表，列表查询一般涉及分页
// func GetEmptyBoxQuantity(pageSize int, pageNum int) ([]EmptyBoxQuantity, int) {
// 	var EmptyBoxQuantitylist []EmptyBoxQuantity
// 	var err error
// 	if pageSize == -1 && pageNum == -1 {
// 		err = db.Preload("Point").Limit(-1).Offset(-1).Find(&EmptyBoxQuantitylist).Error
// 	} else {
// 		err = db.Preload("Point").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&EmptyBoxQuantitylist).Error
// 	}

// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, errmsg.ERROR
// 	}
// 	return EmptyBoxQuantitylist, errmsg.SUCCESS
// }

// // 查询装卸点是否存在
// func CheckEmptyBoxQuantity(title string) int {
// 	var atricle EmptyBoxQuantity
// 	db.Select("id").Where("Title = ?", title).First(&atricle)
// 	if atricle.ID > 0 {
// 		return errmsg.ERROR_CATE_USED
// 	}
// 	return errmsg.SUCCESS
// }

// // // 编辑文章（根据id）
// // func EidtEbq(id int, data *EmptyBoxQuantity) int {
// // 	var EmptyBoxQuantity EmptyBoxQuantity
// // 	var maps = make(map[string]interface{})
// // 	maps["cid"] = data.Cid
// // 	maps["title"] = data.Title
// // 	maps["desc"] = data.Desc
// // 	maps["content"] = data.Content
// // 	maps["img"] = data.Img

// // 	err = db.Model(&EmptyBoxQuantity).Where("id = ?", id).Updates(maps).Error
// // 	if err != nil {
// // 		return errmsg.ERROR
// // 	}
// // 	return errmsg.SUCCESS
// // }

// // 删除文章
// func DelEmptyBoxQuantity(id int) int {
// 	var EmptyBoxQuantity EmptyBoxQuantity
// 	// 软删除
// 	err = db.Where("id = ?", id).Delete(&EmptyBoxQuantity).Error
// 	if err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCESS
// }
