// 入口文件，连接配置数据库

package model

import (
	"autotrans/utils"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDb() {
	// gorm连接到数据库
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPost,
		utils.DbName,
	)
	// fmt.Println(dsn)
	// fmt.Println("Password:", utils.DbPassWord)

	// 使用gorm连接并打开数据库，后续使用db处理数据
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		// DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("数据库连接失败", err)
		os.Exit(1)
	}

	// 传入结构体地址，自动迁移 schema，保持 schema 是最新的。
	db.AutoMigrate(&User{}, &Box{}, &Point{}, &MaterialTransportRecord{}) // 自动迁移

	// 连接池设置
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("数据库连接池设置失败")
	}
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
