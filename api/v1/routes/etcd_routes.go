package routes

import (
	"github.com/ciliverse/cilikube/api/v1/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterEtcdRoutes(router *gin.RouterGroup, handler *handlers.EtcdHandler) {
	etcdRoutes := router.Group("/etcd")
	{
		etcdRoutes.GET("/list", handler.ListKeys)
	}
}
