package routes

import (
	v1 "autotrans/api/v1" // 重命名
	"autotrans/middleware"
	"autotrans/utils"

	"github.com/gin-gonic/gin"
)

// import "net/http"

// 新建路由入口的方法，外部可以访问
func InitRouter() {
	gin.SetMode(utils.AppMode) // 这里的AppMode是utils包初始化时使用goini从config.ini文件中读取到的
	r := gin.Default()

	// 初始化路由，路由入口
	// 分为需要鉴权和公共两组接口
	routerV1 := r.Group("api/v1") // 新建一个路由组（V1是方便版本管理）

	routerV1.Use(middleware.JwtToken()) // 需要jwt验证
	{
		// 绑定路由
		// user模型的路由接口
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DelUser)

		// // ebq的路由接口
		// routerV1.POST("ebq/add", v1.AddEbq)
		// routerV1.GET("ebq", v1.GetEbq)
		// routerV1.PUT("ebq/:id", v1.EditEbq)
		// routerV1.DELETE("ebq/:id", v1.DelEbq)
		// routerV1.GET("ebq/list/:id", v1.GetPointEbq)
		// routerV1.GET("ebq/info/:id", v1.GetEbqInfo)
	}

	publicRouterV1 := r.Group("api/v1") // 新建一个路由组（V1是方便版本管理）
	{

		publicRouterV1.POST("login", v1.Login)
		publicRouterV1.POST("user/add", v1.AddUser)
		publicRouterV1.GET("user", v1.GetUsers)

		// point的路由接口
		publicRouterV1.POST("point/add", v1.AddPoint)
		publicRouterV1.GET("point", v1.GetPoint)
		publicRouterV1.PUT("point/:id", v1.EditPoi)
		publicRouterV1.DELETE("point/:id", v1.DelPoi)

		// box的路由接口
		publicRouterV1.POST("box/addone", v1.AddOneBox)
		publicRouterV1.PUT("box/addmuti/", v1.AddMutiBox)
		publicRouterV1.DELETE("box/:pointid", v1.DelOneBox)
		publicRouterV1.DELETE("box/all/:pointid", v1.DelBoxByPoint)
		publicRouterV1.GET("box/:pointid", v1.GetBox)

		// publicRouterV1.GET("ebq", v1.GetEbq)
		// publicRouterV1.GET("ebq/list/:id", v1.GetPointEbq)
		// publicRouterV1.GET("ebq/info/:id", v1.GetEbqInfo)
	}

	r.Run(utils.HttpPort)
}
