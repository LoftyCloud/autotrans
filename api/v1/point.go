package v1

import (
	"autotrans/model"
	"autotrans/utils/errmsg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 装卸点模型 的控制器接口

// 添加装卸点
func AddPoint(c *gin.Context) {
	var data model.Point
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Bind JSON Error:", err)
	}

	code = model.CheckPoi(data.PointName) // 检查装卸点是否存在
	if code == errmsg.SUCCESS {
		model.CreatePoi(&data) // 添加装卸点
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 获取Point列表
func GetPoint(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetPoint(pageSize, pageNum) // 调用方法
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑Point信息
func EditPoi(c *gin.Context) {
	var data model.Point
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckPoi(data.PointName)
	if code == errmsg.SUCCESS {
		model.EditPoi(id, &data)
	}
	if code == errmsg.ERROR_POINT_EXIST {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除装卸点
func DelPoi(c *gin.Context) {
	// todo: 如果还有空料框则不允许删除

	id, _ := strconv.Atoi(c.Param("id"))
	// fmt.Println(id)
	code = model.DelPoi(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"massage": errmsg.GetErrMsg(code),
	})
}
