package controller


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTaskByID(c *gin.Context){
	c.String(http.StatusOK,"Hello world")
}