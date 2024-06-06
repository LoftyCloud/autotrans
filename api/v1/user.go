package v1

import (
	"fmt"
	"autotrans/model"
	"autotrans/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user模型的【控制器接口】
var code int  // 错误码

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil { // 接收数据
		fmt.Println("Bind JSON Error:", err)
	}
	// fmt.Println(data)
	code = model.CheckUser(data.Username) // 检查用户是否存在
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)  // 调用方法创建用户
	}
	c.JSON(http.StatusOK, gin.H{  // 向前端返回json消息
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表,分页可以在数据量较大时提高性能
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	
	data := model.GetUsers(pageSize, pageNum) // 调用方法
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&data); err != nil { // 接收数据
		fmt.Println("Bind JSON Error:", err)
	}

	fmt.Println(data.Username)
	code = model.CheckUser(data.Username)  // 检查修改后的用户名是否重复
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	code = model.DelUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"massage": errmsg.GetErrMsg(code),
	})
}
