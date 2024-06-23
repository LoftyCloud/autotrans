package v1

import (
	"autotrans/model"
	"autotrans/utils/errmsg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// import (
// 	"fmt"
// 	"autotrans/model"
// 	"autotrans/utils/errmsg"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// 为某个point插入一个分料框
func AddOneBox(c *gin.Context) {
	var data model.Box
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Bind JSON Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  errmsg.ERROR,
			"message": "Invalid JSON data",
		})
		return
	}
	// 检查point是否存在
	code := model.CheckPoiByID(data.PointID)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code = model.CreateBox(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 为某个point插入多个分料框
func AddMutiBox(c *gin.Context) {
	pointID, _ := strconv.Atoi(c.Query("pointid"))
	num, _ := strconv.Atoi(c.Query("num"))

	code = model.CheckPoiByID(pointID)
	if code == errmsg.SUCCESS {
		if num >= 0 {
			model.CreateMutiBox(pointID, num)
		} else {
			code = errmsg.ERROR_BOX_NUM
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除一个分料框
func DelOneBox(c *gin.Context) {
	pointid, _ := strconv.Atoi(c.Param("pointid"))
	code := model.CheckPoiByID(pointid)

	if code == errmsg.SUCCESS {
		model.DeleteBox(pointid)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除某个point下的所有分料框
func DelBoxByPoint(c *gin.Context) {
	pointid, _ := strconv.Atoi(c.Param("pointid"))
	code := model.CheckPoiByID(pointid)
	if code == errmsg.SUCCESS {
		model.DeleteBoxAll(pointid)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询某个point下的所有分料框
func GetBox(c *gin.Context) {
	pointid, _ := strconv.Atoi(c.Param("pointid"))
	code := model.CheckPoiByID(pointid)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	data := model.GetBox(pointid)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
