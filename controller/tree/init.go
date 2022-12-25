package tree

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) []*gin.RouterGroup {
	data := make([]*gin.RouterGroup, 0)
	data = append(data, InitProducts(e))
	data = append(data, InitApps(e))
	data = append(data, InitGroups(e))
	return data
}

func InitProducts(e *gin.Engine) *gin.RouterGroup {
	productApi := e.Group("/api/v1/products")

	productApi.GET("")
	productApi.GET("/:id")
	productApi.POST("")
	productApi.PUT("/:id")
	productApi.DELETE("/:id")

	return productApi
}

func InitApps(e *gin.Engine) *gin.RouterGroup {
	appApi := e.Group("/api/v1/apps")

	appApi.GET("")
	appApi.GET("/:id")
	appApi.POST("")
	appApi.PUT("/:id")
	appApi.DELETE("/:id")

	return appApi
}

func InitGroups(e *gin.Engine) *gin.RouterGroup {
	groupApi := e.Group("/api/v1/groups")

	groupApi.GET("")
	groupApi.GET("/:id")
	groupApi.POST("")
	groupApi.PUT("/:id")
	groupApi.DELETE("/:id")

	return groupApi
}
