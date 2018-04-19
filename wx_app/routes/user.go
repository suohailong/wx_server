package routes

import (
	"github.com/gin-gonic/gin"
	contro "wx_app/controller"
)

func SetUserRoutes(router *gin.Engine) *gin.Engine {
	taR := router.Group("/tm2/")
  taR.GET("/t", contro.GetTaskByID)
	return router
}