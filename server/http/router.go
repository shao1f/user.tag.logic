package http

import (
	"github.com/gin-gonic/gin"
)

func initRoute(router *gin.Engine) {
	v1 := router.Group("/api/user/tag")
	v1.POST("/add", tagAdd)
	v1.GET("/search", tagSearch)
	v1.POST("/link_entity", entityLink)
	v1.GET("/entity_tags", entitySearch)
}
