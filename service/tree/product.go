package tree

import (
	"AppCMDBService/dao"
	"AppCMDBService/logger"
	"AppCMDBService/view"
	"fmt"
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
