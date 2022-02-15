package routes

import (
	"github.com/gin-gonic/gin"
	"lmwn-exam/controllers"
)

func UserRoute(router *gin.Engine)  {
	router.GET("/covid/summary", controllers.Summary())
}