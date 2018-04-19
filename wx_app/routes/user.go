package routes

import (
	"github.com/gin-gonic/gin"
	contro "controller"
)

func SetUserRoutes(router *gin.Engine) *gin.Engine {
	taR := router.Group("/tm2/")
  taR.GET("/t", contro.GetTaskByID)
	return router
}