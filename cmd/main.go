package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRouters(router)

	router.Run()

}
