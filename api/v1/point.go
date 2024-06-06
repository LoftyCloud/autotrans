package v1

// import (
// 	"fmt"
// 	"autotrans/model"
// 	"autotrans/utils/errmsg"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// 装卸点模型 的控制器接口

// // 获取Point列表
// func GetPoint(c *gin.Context) {
// 	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
// 	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

// 	if pageSize == 0 {
// 		pageSize = -1
// 	}
// 	if pageNum == 0 {
// 		pageNum = -1
// 	}

// 	data := model.GetPoint(pageSize, pageNum)  // 调用方法
// 	code = errmsg.SUCCESS
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  code,
// 		"data":    data,
// 		"message": errmsg.GetErrMsg(code),
// 	})
// }

// // 添加装卸点
// func AddPoint(c *gin.Context) {
// 	var data model.Point
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		fmt.Println("Bind JSON Error:", err)
// 	}

// 	fmt.Println(data)

// 	code = model.CheckCate(data.Name)
// 	if code == errmsg.SUCCESS {
// 		model.CreateCate(&data)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  code,
// 		"data":    data,
// 		"message": errmsg.GetErrMsg(code),
// 	})
// }

// // todo 查询单个分类下的文章

// // 编辑分类名
// func EditCate(c *gin.Context) {
// 	var data model.Category
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	c.ShouldBindJSON(&data)
// 	code = model.CheckCate(data.Name)
// 	if code == errmsg.SUCCESS {
// 		model.EidtCate(id, &data)
// 	}
// 	if code == errmsg.ERROR_CATE_USED {
// 		c.Abort()
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  code,
// 		"message": errmsg.GetErrMsg(code),
// 	})
// }

// // 删除装卸点
// func DelCate(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	fmt.Println(id)
// 	code = model.DelCate(id)
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  code,
// 		"massage": errmsg.GetErrMsg(code),
// 	})
// }
