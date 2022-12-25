package view

import (
	"AppCMDBService/model"
)

type ProductListRequest struct {
	ProductName string `json:"product_name"` // 名称
	Code        string `json:"code"`         // 代码
	ProductType string `json:"product_type"` // 类型
	IsDelete    uint16 `json:"is_delete"`    // 是否删除
	CreateUser  string `json:"create_user"`  // 创建人
	ParentId    int64  `json:"parent_id"`    // 父节点ID，t_product.id

	Search    string `json:"search"`
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
}

type ProductListResponse struct {
	Count    uint64           `json:"count"`
	Records  []*model.Product `json:"records"`
	Page     int64            `json:"page"`
	PageSize int64            `json:"page_size"`
}
