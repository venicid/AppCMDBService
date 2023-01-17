package tree

import (
	"AppCMDBService/dao/mysql"
	"AppCMDBService/logger"
	"AppCMDBService/setting"
	"AppCMDBService/view"
	"fmt"
	"go.uber.org/zap"
	"testing"
)

func InitTools() {
	setting.InitConfig("E:\\helloGolang\\src\\AppCMDBService\\config\\config.yaml")
	logger.Init()
	defer zap.L().Sync()
	mysql.Init()
}

func TestListProductRecords(t *testing.T) {
	InitTools()

	req := &view.ProductListRequest{}

	records, err := ListProductRecords(req)
	if err != nil {
		fmt.Println("failed")
		return
	}

	for _, v := range records.Records {
		fmt.Println(v)
	}

}

func TestCreateProductRecord2(t *testing.T) {
	InitTools()

	p := 9
	req := &view.ProductCreateRequest{
		ProductName: "基础技术部",
		Code:        "BU-BASETECH",
		ProductType: "team",
		CreateUser:  "alex",
		ParentId:    &p,
	}

	CreateProductRecord(req)

}

func TestCreateBatchProductRecord(t *testing.T) {
	InitTools()

	group := map[string]string{
		"CMDB系统":  "SysCMDB",
		"发布系统":    "SysCICD",
		"工单系统":    "SysTicket",
		"自动化运维系统": "SysAutoOps",
	}

	p := 38
	req2 := &view.ProductCreateRequest{
		ProductType: "team",
		CreateUser:  "alex",
		ParentId:    &p,
	}

	for k, v := range group {
		req2.ProductName = k
		req2.Code = v
		CreateProductRecord(req2)
	}

	group1 := map[string]string{
		"中国商业":    "ChinaCommerce",
		"国际商业":    "InternationalCommerce",
		"本地生活服务":  "LocalConsumerServices",
		"数字媒体及娱乐": "DigitalMediaAndEntertainment",
		"创新业务及其他": "InnovationInitiativesAndOthers",
		"物流":      "Logistics",
		"云":       "Cloud",
	}
	groupParentId := 1
	req1 := &view.ProductCreateRequest{
		ProductType: "team",
		CreateUser:  "mayun002",
		ParentId:    &groupParentId,
	}

	for k, v := range group1 {
		req1.ProductName = k
		req1.Code = v
		//CreateProductRecord(req1)
	}
}
