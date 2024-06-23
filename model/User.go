package model

import (
	"autotrans/utils/errmsg"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

// 【user模型与具体操作方法】
// UserID	INT	用户ID（主键）
// Username	VARCHAR	用户名
// Password	VARCHAR	密码（加密存储）
// Email	VARCHAR	电子邮箱
// Role	VARCHAR	用户角色（例如管理员、操作员等）
type User struct {
	gorm.Model // 引入gorm中定义的变量，包括ID以及创建时间等
	// 绑定json格式，方便与前端交互
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Email    string `gorm:"type:varchar(20);" json:"email"`
	Role     string `gorm:"varchar(20);not null" json:"role"`
}

// // 方法的具体实现,在接口处被调用,这里主要实现了对数据库的操作
// 查询用户是否存在
func CheckUser(userName string) int {
	var user User
	db.Select("id").Where("username = ?", userName).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 结构体为引入变量，传入方法时使用指针
func CreateUser(data *User) int {
	// 密码写入数据库前进行加密
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密存储
func ScryptPw(password string) string {
	const KeyLen = 10
	// salt := make([]byte, 8)
	salt := []byte{1, 5, 3, 65, 2, 42, 32, 4}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	SHashPw := base64.StdEncoding.EncodeToString(HashPw)
	return SHashPw
}

// 查询用户列表，列表查询一般涉及分页
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	var err error
	if pageSize == -1 && pageNum == -1 {
		err = db.Limit(-1).Offset(-1).Find(&users).Error
	} else {
		err = db.Select("username", "id").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("GetUsers ERROR")
		return nil
	}
	return users
}

// 根据用户id编辑用户信息（password和role除外）
func EditUser(id int, data *User) int {
	// var user User
	var maps = make(map[string]interface{})
	// 使用Map传参,若使用结构体传参,则参数为0的部分不会更新,参看gorm手册.
	maps["username"] = data.Username
	maps["email"] = data.Email
	err = db.Model(User{}).Where("id = ?", id).Updates(maps).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DelUser(id int) int {
	var user User
	// 软删除

	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		// fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 登录验证
func CheckLogin(userName string, password string) int {
	var user User
	db.Where("username=?", userName).First(&user)

	// 用户不存在
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	// 可在此进行权限管理，如普通用户无法登录后台等功能，对user.Role进行判断
	return errmsg.SUCCESS
}
