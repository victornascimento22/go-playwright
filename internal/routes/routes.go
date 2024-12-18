package routes

import (
"gitlab.com/applications2285147/api-go/internal"

)
func RegisterRoutes(router *gin.Engine){

	router.GET("/health", func(c *gin.Context){
		c.JSON(200, gin.H("status":"ok"))
	})
}

