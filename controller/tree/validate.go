package tree

import (
	"AppCMDBService/controller/common"
	"AppCMDBService/view"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getProductListParams(c *gin.Context) (params *view.ProductListRequest, err error) {
	params = &view.ProductListRequest{
		ProductName: c.DefaultQuery("product_name", ""),
		Code:        c.DefaultQuery("code", ""),
		ProductType: c.DefaultQuery("product_type", ""),
		CreateUser:  c.DefaultQuery("create_user", ""),
		Search:      c.DefaultQuery("search", ""),
	}

	params.Page, params.PageSize = common.GetPagePageSizeParams(c)
	params.SortField, params.SortOrder = common.GetSortFieldDesc(c)

	params.IsDelete, err = common.GetIsDeleteParams(c)
	if err != nil {
		return nil, err
	}

	parentIdStr := c.Query("parent_id")
	if parentIdStr != "" {
		parentId, err := strconv.Atoi(parentIdStr)
		if err != nil {
			return nil, err
		}
		params.ParentId = int64(parentId)
	}

	return params, err
}

func getProductCreateParams(c *gin.Context) (params *view.ProductCreateRequest, err error) {
	params = &view.ProductCreateRequest{}
	err = c.ShouldBind(params)
	if err != nil {
		return nil, errors.New("bind绑定参数失败")
	}

	if params.ProductType == "" {
		return nil, errors.New("ProductType不能为空")
	}
	if params.ProductName == "" {
		return nil, errors.New("ProductName不能为空")
	}
	if params.Code == "" {
		return nil, errors.New("code不能为空")
	}
	if params.CreateUser == "" {
		return nil, errors.New("CreateUser不能为空")
	}
	if params.ParentId == nil {
		return nil, errors.New("ParentId不能为空")
	}
	return params, err

}
