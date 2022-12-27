package tree

import (
	"AppCMDBService/dao"
	"AppCMDBService/logger"
	"AppCMDBService/model"
	"AppCMDBService/view"
	"fmt"
	"time"
)

func ListProductRecords(req *view.ProductListRequest) (*view.ProductListResponse, error) {

	var response = &view.ProductListResponse{}

	err, result, count := dao.ListProductRecords(req)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("[service.ListProductRecords], result: %v, count:%d, request: %v", result, count, req)
	logger.Logger.Warn(msg)

	response.Count = count
	response.Records = result
	if req.Page == -1 {
		response.Page = 1
		response.PageSize = int64(count)
	} else {
		response.Page = req.Page
		response.PageSize = req.PageSize
	}

	return response, nil
}

func GetProductDetail(productId int) (*model.Product, error) {

	record, err := dao.GetProductRecord(productId)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("[service.GetProductDetail], result: %v,  productId: %v", record, productId)
	logger.Logger.Warn(msg)

	return record, nil
}

func GetProductByName(productName string) (*model.Product, error) {

	record, err := dao.GetProductRecordByName(productName)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("[service.GetProductByName], result: %v,  productName: %v", record, productName)
	logger.Logger.Warn(msg)

	return record, nil
}

func CreateProductRecord(params *view.ProductCreateRequest) (err error) {
	product := &model.Product{
		ProductName: params.ProductName,
		Code:        params.Code,
		ProductType: params.ProductType,
		IsDelete:    0,
		CreateUser:  params.CreateUser,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		ParentId:    *params.ParentId,
	}

	err = dao.CreateProductRecord(product)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("[service.CreateProductRecord],  product: %v", product)
	logger.Logger.Warn(msg)

	return nil
}
