package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	{
		// Add /hello GET route to router and define route handler function
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "toast"})
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"status": "error"}) })

	// Start listening and serving requests
	router.Run(":8080")
}
