package tree

import (
	"AppCMDBService/controller/common"
	"AppCMDBService/logger"
	"AppCMDBService/service/tree"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetTreeGlobalHandler(c *gin.Context) {
	result, err := tree.GetProductTree()
	if err != nil {
		msg := fmt.Sprintf("GetProductListHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, "获取产品树列表错误", nil)
		return
	}

	common.HttpResponse(c, 200, "成功", result)
}
