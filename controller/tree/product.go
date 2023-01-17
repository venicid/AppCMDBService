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

func GetMiniProductRecords(c *gin.Context) {

	result, err := tree.ListMiniProductRecords()
	if err != nil {
		msg := fmt.Sprintf("GetMiniProductRecords.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		c.JSON(400, gin.H{"code": 400, "msg": "获取mini-product列表错误", "data": nil})
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

func CreateProductHandler(c *gin.Context) {
	params, err := getProductCreateParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("CreateProductHandler.error.message.%s", err.Error()))
		c.JSON(400, gin.H{"code": 400, "msg": err.Error(), "data": nil})
		return
	}

	// 判断已经存在
	record, err := tree.GetProductByName(params.ProductName)
	if err != nil && record != nil {
		logger.Logger.Error(fmt.Sprintf("CreateProductHandler.ListProductRecords.error.message.%s", err.Error()))
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("查询product错误， err:%s", err.Error()), "data": nil})
		return
	}
	if record != nil {
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("该product已存在, productId:%v", record.Id), "data": nil})
		return
	}

	err = tree.CreateProductRecord(params)
	if err != nil {
		msg := fmt.Sprintf("CreateProductHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("创建product错误， err:%s", err.Error()), "data": nil})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "成功", "data": nil})
}

func UpdateProductHandler(c *gin.Context) {
	params, err := getProductUpdateParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("UpdateProductHandler.error.message.%s", err.Error()))
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	// id参数
	id, err := common.GetIdParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("GetProductDetailHandler.error.message.%s", err.Error()))
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	// 判断是否存在
	record, err := tree.GetProductDetail(id)
	if err != nil || record == nil {
		logger.Logger.Error(fmt.Sprintf("UpdateProductHandler.GetProductDetail.error.message.%s", err.Error()))
		common.HttpResponse(c, 400, "该product不存在, 无法更新", nil)
		return
	}

	// 是否需要更新
	updateFlag := false
	if params.ProductName != "" && params.ProductName != record.ProductName {
		record.ProductName = params.ProductName
		updateFlag = true
	}
	if params.Code != "" && params.Code != record.Code {
		record.Code = params.Code
		updateFlag = true
	}
	if params.ProductType != "" && params.ProductType != record.ProductType {
		record.ProductType = params.ProductType
		updateFlag = true
	}
	if params.CreateUser != "" && params.CreateUser != record.CreateUser {
		record.CreateUser = params.CreateUser
		updateFlag = true
	}
	if params.ParentId != nil && *params.ParentId != record.ParentId {
		record.ParentId = *params.ParentId
		updateFlag = true
	}
	if params.IsDelete != nil && *params.IsDelete != record.IsDelete {
		record.IsDelete = *params.IsDelete
		updateFlag = true
	}

	if !updateFlag {
		common.HttpResponse(c, 200, "该记录未改变", nil)
		return
	}

	err = tree.UpdateProductRecord(record)
	if err != nil {
		msg := fmt.Sprintf("UpdateProductHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		c.JSON(400, gin.H{"code": 400, "msg": fmt.Sprintf("更新product错误， err:%s", err.Error()), "data": nil})
		return
	}

	c.JSON(200, gin.H{"code": 200, "msg": "成功", "data": nil})
}

func DeleteSoftProductHandler(c *gin.Context) {
	// id参数
	id, err := common.GetIdParams(c)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("GetProductDetailHandler.error.message.%s", err.Error()))
		common.HttpResponse(c, 400, err.Error(), nil)
		return
	}

	// 判断是否存在
	record, err := tree.GetProductDetail(id)
	if err != nil || record == nil {
		logger.Logger.Error(fmt.Sprintf("UpdateProductHandler.GetProductDetail.error.message.%s", err.Error()))
		common.HttpResponse(c, 400, "该product不存在, 无法更新", nil)
		return
	}

	err = tree.DeleteSoftProductRecord(record)
	if err != nil {
		msg := fmt.Sprintf("DeleteSoftProductHandler.error.message.%s", err.Error())
		logger.Logger.Error(msg)
		common.HttpResponse(c, 400, fmt.Sprintf("删除product错误， err:%s", err.Error()), nil)

		return
	}

	common.HttpResponse(c, 200, "成功", nil)
}
