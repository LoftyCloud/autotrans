package errmsg

// 错误处理模块，声明常量错误码
const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000* 用户模块
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003

	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	
	// // code = 2000* 模块
	// ERROR_Art_NOT_EXIST = 2001
	// // code = 3000* 分类模块
	// ERROR_CATE_USED      = 3001
	// ERROR_CATE_NOT_EXIST = 3002
)

// 错误消息字典
var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",

	// ERROR_CATE_USED:      "分类存在",
	// ERROR_CATE_NOT_EXIST: "分类不存在",

	// ERROR_Art_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}