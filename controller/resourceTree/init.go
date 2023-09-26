package resourceTree

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)

	data = append(data, InitProducts(e))
	data = append(data, InitTree(e))
	data = append(data, InitApps(e))
	//data = append(data, InitGroups(e))
	return data
}

//
//func InitGroups(e *gin.Engine) *gin.RouterGroup {
//	groupApi := e.Group("/api/v1/groups")
//
//	groupApi.GET("")
//	groupApi.GET("/:id")
//	groupApi.POST("")
//	groupApi.PUT("/:id")
//	groupApi.DELETE("/:id")
//
//	return groupApi
//}
