package routes

import (
	"github.com/gin-gonic/gin"
	contro "../controller"
)

func SetUserRoutes(router *gin.Engine) *gin.Engine {
	taR := router.Group("/wx/")
  taR.GET("/t", contro.VerifyWx)
	return router
}