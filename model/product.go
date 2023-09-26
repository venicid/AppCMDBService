package model

import (
	"time"
)

type Product struct {
	Id          int       `gorm:"column:id;primary_key" json:"id"`                  // 主键
	ProductName string    `gorm:"column:product_name;NOT NULL" json:"product_name"` // 名称
	Code        string    `gorm:"column:code;NOT NULL" json:"code"`                 // 代码
	ProductType string    `gorm:"column:product_type;NOT NULL" json:"product_type"` // 类型
	IsDelete    int       `gorm:"column:is_delete;NOT NULL" json:"is_delete"`       // 是否删除
	CreateUser  string    `gorm:"column:create_user;NOT NULL" json:"create_user"`   // 创建人
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`            // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`            // 更新时间
	ParentId    int       `gorm:"column:parent_id" json:"parent_id"`                // 父节点ID，t_product.id
}

type ProductMini struct {
	Id          *int   `gorm:"column:id;primary_key" json:"id"`                  // 主键
	ProductName string `gorm:"column:product_name;NOT NULL" json:"product_name"` // 名称
	Code        string `gorm:"column:code;NOT NULL" json:"code"`                 // 代码
	ProductType string `gorm:"column:product_type;NOT NULL" json:"product_type"` // 类型
	ParentId    *int   `gorm:"column:parent_id" json:"parent_id"`                // 父节点ID，t_product.id
}

func (m *Product) TableName() string {
	return "t_product"
}
