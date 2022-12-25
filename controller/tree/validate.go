package tree

import (
	"AppCMDBService/controller/common"
	"AppCMDBService/view"
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
