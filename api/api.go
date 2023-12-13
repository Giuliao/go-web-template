package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"my.net/module/pkg/handler"
)

// Register api
func Register(r *gin.Engine) {

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/ping", handler.Ping)

	apiV2 := r.Group("/api/v2")
	apiV2.GET("/ping", handler.PPing)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
