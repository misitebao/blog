package router

import (
	"blog/app/controller"
	"github.com/gin-gonic/gin"
)

// 初始化函数
func RegisteRouter() {
	r := gin.Default()
	//加载模板目录
	r.LoadHTMLGlob("**/view/*")
	r.Use(RunTime)
	//根路由
	r.GET("/", controller.PageController{}.Index)

	//第一版API组
	v1 := r.Group("/v1")
	{
		//用户
		v1.GET("/info", controller.UserController{}.Login)
	}
	//第二版API组
	v2 := r.Group("/v2")
	{
		//用户
		v2.GET("/info", controller.UserController{}.Login)
	}

	r.Run()
}
