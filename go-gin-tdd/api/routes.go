package api

import (
	"github.com/alochym01/ftp-api/api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/get", controllers.GetPing)
	router.POST("/pingjson", controllers.PostPingJson)
	router.POST("/pingform", controllers.PostPingForm)
	return router
}
