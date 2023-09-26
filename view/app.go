package view

import "AppCMDBService/model"

type AppListRequest struct {
	AppName  string `json:"app_name"` // 名称
	AppLevel string `json:"app_level"`
	Lang     string `json:"lang"`

	ProductId  string `json:"product_id"`
	IsDelete   uint16 `json:"is_delete"`   // 是否删除
	CreateUser string `json:"create_user"` // 创建人

	Search    string `json:"search"`
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	Page      int64  `json:"page"`
	PageSize  int64  `json:"page_size"`
}

type AppListResponse struct {
	Count    uint64       `json:"count"`
	Records  []*model.App `json:"records"`
	Page     int64        `json:"page"`
	PageSize int64        `json:"page_size"`
}
