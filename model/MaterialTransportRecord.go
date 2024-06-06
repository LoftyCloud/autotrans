package model

import (
	// "autotrans/utils/errmsg"
	"gorm.io/gorm"
)

// 物料运输记录表
// RecordID	INT	记录ID（主键）
// TransportDateTime	DATETIME	运输时间
// MaterialType	VARCHAR	物料类型
// TransportedBy	INT	运输人员（外键，关联User表的UserID）
type MaterialTransportRecord struct {
	User User `gorm:"foreignkey:TransportedBy"`  // 外键，关联User表
	gorm.Model
	TransportedBy   int    `gorm:"type:int;not null" json:"transportedby"`
	MaterialType  string `gorm:"type:int" json:"materialtype"`
}