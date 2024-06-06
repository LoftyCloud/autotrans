package v1

// 登录验证模型的控制器接口
import(
	"autotrans/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"autotrans/utils/errmsg"
	"autotrans/middleware"
)

func Login(c *gin.Context){
	var data model.User
	var token string
	var code int
	c.ShouldBindJSON(&data)
	// 检查账号密码是否正确
	code = model.CheckLogin(data.Username,data.Password)
	
	// token验证
	if code == errmsg.SUCCESS{ // 设置token
		token,code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"token":token,
	})
}