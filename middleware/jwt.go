package middleware

import (
	"autotrans/utils"
	"autotrans/utils/errmsg"
	"time"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(utils.JwtKey)  // 密钥

type MyCLaims struct {  // 验证用信息；用户名。时间，签发人
	Username string `json:"username"`
	jwt.StandardClaims  // jwt设置的标准结构体
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)  // 区间有效期
	SetClaim := MyCLaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //有效时间
			Issuer:    "aurotrans-user", // 签发人
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaim) //接收 签发方法+声明，返回一个token
	token, err := reqClaim.SignedString(JwtKey)  // 返回签名token
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// 验证token
func CheckToken(token string) (*MyCLaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyCLaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyCLaims); setToken.Valid {
		return key, errmsg.SUCCESS
	}
	return nil, errmsg.ERROR
}

// jwt中间件(固定写法，参看jwt文档)
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {  // 返回一个函数
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader," ")
		if len(checkToken)!=2 && checkToken[0]!="Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR{
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix()>key.ExpiresAt{
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
