package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routers
	initializeRouters(router) 
	// Run the server
	router.Run(":8080")

}
