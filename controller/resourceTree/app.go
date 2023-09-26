package resourceTree

import (
	"AppCMDBService/service/resourceTree"
	"github.com/gin-gonic/gin"
)

func InitApps(e *gin.Engine) *gin.RouterGroup {
	appApi := e.Group("/api/v1/apps")

	appApi.GET("", GetAppListHandler)
	appApi.GET("/:id", GetAppDetailHandler)
	appApi.POST("", CreateAppHandler)
	appApi.PUT("/:id", UpdateAppHandler)
	appApi.DELETE("/:id", DeleteSoftAppHandler)

	return appApi
}

func GetAppListHandler(c *gin.Context) {

	resourceTree.ListAppRecords(nil)
}

func GetAppDetailHandler(c *gin.Context) {
	resourceTree.GetAppDetail()

}

func CreateAppHandler(c *gin.Context) {
	resourceTree.CreateAppRecord()

}

func UpdateAppHandler(c *gin.Context) {
	resourceTree.UpdateAppRecord()

}

func DeleteSoftAppHandler(c *gin.Context) {
	resourceTree.DeleteSoftAppRecord()

}
