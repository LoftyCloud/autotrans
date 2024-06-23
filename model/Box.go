package model

import (
	"autotrans/utils/errmsg"
	"fmt"

	"gorm.io/gorm"
)

// 空料框数量记录表
// RecordID	INT	记录ID（主键）
// PointID	INT	装卸点ID（外键，关联装卸点表）
// DateTime	DATETIME	记录时间
type Box struct {
	Point Point `gorm:"foreignkey:PointID"` // 外键
	gorm.Model
	PointID int  `gorm:"type:int;not null" json:"pointid"`
	Empty   bool `gorm:"type:bool;not null" json:"empty"`
}

// 方法实现
// 新增一个框，结构体为引入变量，传入方法时使用指针
func CreateBox(data *Box) int {
	err := db.Create(data).Error
	if err != nil {
		fmt.Println("Create Error:", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 为某个point插入多个分料框
func CreateMutiBox(pointID int, num int) int {
	for i := num; i > 0; i-- { // 循环插入数据
		var box = Box{
			PointID: pointID,
			Empty:   true,
		}

		err := db.Create(&box).Error
		if err != nil {
			return errmsg.ERROR
		}
	}

	return errmsg.SUCCESS
}

// 通过PointID删除一个分料框
func DeleteBox(pointid int) int {
	var box Box
	err := db.Where("point_id=?", pointid).First(&box).Error
	if err != nil {
		return errmsg.ERROR
	}

	err = db.Delete(&box).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除某个point下的所有分料框
func DeleteBoxAll(pointid int) int {
	var box Box
	// 删除point_id为pointid的box记录
	err := db.Where("point_id = ?", pointid).Delete(&box).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询某个point下的所有分料框
func GetBox(pointid int) []Box {
	var box []Box
	err := db.Where("point_id=?", pointid).Find(&box).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return box
}
