package model

import (
	"time"
)

type App struct {
	Id                 int       `gorm:"column:id;primary_key" json:"id"`
	AppName            string    `gorm:"column:app_name;NOT NULL" json:"app_name"`   // 名称
	AppLevel           int       `gorm:"column:app_level;NOT NULL" json:"app_level"` // 等级
	Lang               string    `gorm:"column:lang;NOT NULL" json:"lang"`           // 类型
	Attributes         string    `gorm:"column:attributes" json:"attributes"`        // 补充参数
	Remark             string    `gorm:"column:remark" json:"remark"`                // 备注
	IsDelete           int       `gorm:"column:is_delete;NOT NULL" json:"is_delete"` // 是否删除
	CreateUser         string    `gorm:"column:create_user;NOT NULL" json:"create_user"`
	CreateTime         time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime         time.Time `gorm:"column:update_time" json:"update_time"`
	ProductId          int       `gorm:"column:product_id;NOT NULL" json:"product_id"`            // Product ID, t_product.id
	OldAppId           int       `gorm:"column:old_app_id" json:"old_app_id"`                     // 老版app的ID
	OutInternetReason  string    `gorm:"column:out_internet_reason" json:"out_internet_reason"`   // 出公网原因
	IsOutInternet      int       `gorm:"column:is_out_internet" json:"is_out_internet"`           // 是否出公网
	ProSpecificationId int       `gorm:"column:pro_specification_id" json:"pro_specification_id"` // 生产环境规格族ID
	Namespace          string    `gorm:"column:namespace" json:"namespace"`                       // 命名空间名称
	ProOsVersion       string    `gorm:"column:pro_os_version" json:"pro_os_version"`             // 生产环境操作系统版本
	InternetMethod     string    `gorm:"column:internet_method" json:"internet_method"`           // 出公网模式：EIP|NAT
	TeamCode           string    `gorm:"column:team_code" json:"team_code"`                       // app所属team信息，冗余字段，方便查询
}

func (m *App) TableName() string {
	return "`t_app`"
}
