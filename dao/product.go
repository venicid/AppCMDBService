package dao

import (
	"AppCMDBService/dao/mysql"
	"AppCMDBService/logger"
	"AppCMDBService/model"
	"AppCMDBService/view"
	"fmt"
)

func ListProductRecords(req *view.ProductListRequest) (error, []*model.Product, uint64) {
	var queryDb = mysql.GORM

	var records []*model.Product
	var count uint64

	if req.ProductName != "" {
		queryDb = queryDb.Where("product_name = ?", req.ProductName)
	}
	if req.Code != "" {
		queryDb = queryDb.Where("code = ?", req.Code)
	}
	if req.ProductType != "" {
		queryDb = queryDb.Where("product_type = ?", req.ProductType)
	}
	if req.CreateUser != "" {
		queryDb = queryDb.Where("create_user = ?", req.CreateUser)
	}
	if req.Search != "" {
		queryDb = queryDb.Where("product_name like ?", "%"+req.Search+"%")
	}

	if req.Page > 0 && req.PageSize > 0 {
		queryDb = queryDb.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize)
	}

	orderBy := "id desc"
	if req.SortField != "" && req.SortOrder != "" {
		orderBy = fmt.Sprintf("%s %s", req.SortField, req.SortOrder)
	}

	queryDb = queryDb.Where("is_delete = ?", req.IsDelete)
	err := queryDb.Order(orderBy).Find(&records).Count(&count).Error
	if err != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.ListProductRecords sql execute err, err message is %s, req: %v", err, req)
		logger.Logger.Error(msg)
		return err, nil, 0
	}
	return err, records, count
}

func ListMiniProductRecords() (error, []*model.ProductMini) {
	var records []*model.ProductMini

	//err := mysql.GORM.Raw("select id, product_name, code, product_type, parent_id from t_product").Scan(&records).Error
	err := mysql.GORM.Model(model.Product{}).Select([]string{"id", "product_name",
		"code", "product_type", "parent_id"}).Scan(&records).Error
	if err != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.ListMiniProductRecords sql execute err, err message is %s, req: %v", err)
		logger.Logger.Error(msg)
		return err, nil
	}
	return err, records
}

func GetProductRecord(productId int) (record *model.Product, err error) {
	record = &model.Product{}
	tx := mysql.GORM.Where("id = ?", productId).First(&record)
	if tx.Error != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.GetProductRecord.sql.execute.err.message.%s.id.%v", tx.Error, productId)
		logger.Logger.Error(msg)
		return nil, tx.Error
	}

	return record, nil
}

func GetProductRecordByName(productName string) (record *model.Product, err error) {
	record = &model.Product{}
	tx := mysql.GORM.Where("product_name = ?", productName).First(&record)
	if tx.Error != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.GetProductRecordByName.sql.execute.err.message.%s.id.%v", tx.Error, productName)
		logger.Logger.Error(msg)
		return nil, tx.Error
	}

	return record, nil
}

func CreateProductRecord(product *model.Product) (err error) {
	tx := mysql.GORM.Create(&product)
	if tx.Error != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.CreateProductRecord.sql.execute.err.message.%s.product.%v", tx.Error, product)
		logger.Logger.Error(msg)
		return tx.Error
	}
	return nil
}

func UpdateProductRecord(product *model.Product) (err error) {
	// update 不会更新非零值
	// save会更新所有字段
	//tx := mysql.GORM.Model(&model.Product{}).Where("id = ?", product.Id).Updates(&product)
	// https://blog.csdn.net/s2603898260/article/details/122387028
	tx := mysql.GORM.Save(&product)
	if tx.Error != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.UpdateProductRecord.sql.execute.err.message.%s.product.%v", tx.Error, product)
		logger.Logger.Error(msg)
		return tx.Error
	}
	return nil
}

func DeleteSoftProductRecord(product *model.Product) (err error) {
	record := &model.Product{
		IsDelete:   product.IsDelete,
		UpdateTime: product.UpdateTime,
	}
	tx := mysql.GORM.Model(&model.Product{}).Where("id = ?", product.Id).Updates(record)
	if tx.Error != nil {
		msg := fmt.Sprintf("【DB.LOG】 dao.product.UpdateProductRecord.sql.execute.err.message.%s.product.%v", tx.Error, product)
		logger.Logger.Error(msg)
		return tx.Error
	}
	return nil
}
