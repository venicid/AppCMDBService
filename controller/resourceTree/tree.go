package resourceTree

import (
	"AppCMDBService/controller/common"
	"AppCMDBService/logger"
	"AppCMDBService/service/resourceTree"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitTree(e *gin.Engine) *gin.RouterGroup {
	treeApi := e.Group("/api/v1/tree")
	treeApi.GET("/global", GetGlobalTreeHandler)
	return treeApi
}

func GetGlobalTreeHandler(c *gin.Context) {

	result, err := resourceTree.GetGlobalTree()
	if err != nil {
		msg := fmt.Sprintf("GetProductListHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, "获取全局产品树错误", nil)
		return
	}

	common.HttpResponse(c, 200, "成功", result)
}
