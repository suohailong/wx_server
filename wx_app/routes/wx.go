package routes

import (
	contro "../controller"
	"github.com/gin-gonic/gin"
	// contro "wx_app/controller"
)

func SetUserRoutes(router *gin.Engine) *gin.Engine {
	taR := router.Group("/wx/")
  taR.GET("/t", contro.VerifyWx)
	return router
}