package main

import (
	"github.com/gin-gonic/gin"
	"lmwn-exam/routes"
)

func main() {
	router := gin.Default()
	routes.UserRoute(router)
	router.Run("localhost:8080") 
}