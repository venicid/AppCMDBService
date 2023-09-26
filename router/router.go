package router

import (
	"AppCMDBService/controller"
	"AppCMDBService/controller/resourceTree"
	"AppCMDBService/middleware"
	"go.uber.org/zap"

	"AppCMDBService/setting"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(setting.Conf.Mode)

	r := gin.New()

	// 注册两个自定义的中间件
	r.Use(middleware.Cors())
	r.Use(middleware.GinLogger(zap.L()))
	r.Use(middleware.GinRecovery(zap.L(), false))

	// 基础路由拆分
	r.GET("/ping", controller.PingHandler)

	// 路由拆分为不同模块
	resourceTree.Init(r)

	return r
}
