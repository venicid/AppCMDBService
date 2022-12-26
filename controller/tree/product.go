package tree

import (
	"AppCMDBService/controller/common"
	"AppCMDBService/logger"
	"AppCMDBService/service/tree"

	"fmt"
	"github.com/gin-gonic/gin"
)

/**
1. 参数处理
2. 业务逻辑
3. 返回响应
*/
func GetProductListHandler(c *gin.Context) {

	params, err := getProductListParams(c)
	if err != nil {
		msg := fmt.Sprintf("GetProductListHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		//c.JSON(400, gin.H{"code": 1, "msg": "参数解析错误", "data": nil})
		common.HttpResponse(c, 400, "参数解析错误", nil)
	}

	result, err := tree.ListProductRecords(params)
	if err != nil {
		msg := fmt.Sprintf("GetProductListHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		c.JSON(400, gin.H{"code": 400, "msg": "获取product列表错误", "data": nil})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "成功", "data": result})
}

func GetProductDetailHandler(c *gin.Context) {
	id, err := common.GetIdParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("GetProductDetailHandler.error.message.%s", err.Error()))
		c.JSON(400, gin.H{"code": 400, "msg": err.Error(), "data": nil})
		return
	}

	result, err := tree.GetProductDetail(id)
	if err != nil {
		msg := fmt.Sprintf("GetProductListHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("获取product详情错误，id:%v, err:%s", id, err.Error()), "data": nil})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "成功", "data": result})
}
