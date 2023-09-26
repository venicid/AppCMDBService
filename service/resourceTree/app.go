package resourceTree

import (
	"AppCMDBService/dao"
	"AppCMDBService/view"
)

func ListAppRecords(request *view.AppListRequest) (*view.AppListResponse, error) {

	dao.ListAppRecords(request)

	return nil, nil

}

func GetAppDetail() {

	dao.GetAppDetail()

}

func CreateAppRecord() {
	dao.CreateAppRecord()
}

func UpdateAppRecord() {
	dao.UpdateAppRecord()
}

func DeleteSoftAppRecord() {
	dao.DeleteSoftAppRecord()
}
