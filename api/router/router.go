package api

import (
	"github.com/gin-gonic/gin"
)

func Router() {

	r := gin.Default()

	aniversario := r.Group("/aniversario")
	{
		aniversario.GET("/getAniversario", controller.aniversarioController)
	}

}
