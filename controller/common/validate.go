package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPagePageSizeParams(c *gin.Context) (pageUint32 int64, pageSizeUint32 int64) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")
	pageInt32, _ := strconv.Atoi(page)
	pageSizeInt32, _ := strconv.Atoi(pageSize)
	return int64(pageInt32), int64(pageSizeInt32)
}

func GetSortFieldDesc(c *gin.Context) (sortField string, sortOrder string) {
	sortField = c.DefaultQuery("sort_field", "id")
	sortOrder = c.DefaultQuery("sort_order", "asc")
	return sortField, sortOrder
}

func GetIdParams(c *gin.Context) (id int, err error) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, errors.New("id不能为空")
	}
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetIsDeleteParams(c *gin.Context) (id uint16, err error) {
	isDeleteStr := c.DefaultQuery("is_delete", "0")
	isDelete, err := strconv.Atoi(isDeleteStr)
	if err != nil {
		return 0, err
	}
	return uint16(isDelete), nil
}
