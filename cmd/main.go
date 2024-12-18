package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/applications2285147/api-go/internal/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run()

}
