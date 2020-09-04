package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/wmsx/group_api/handler"
)

/**
 * 初始化路由信息
 */
func InitRouter(c client.Client) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	groupHandler := handler.NewGroupHandler(c)
	groupRouter := r.Group("/group")

	groupRouter.POST("/discuss/list", groupHandler.GetAllDiscussGroupList)

	return r
}
