// api/router.go
package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

func Router(db *sql.DB) {
	r := gin.Default()

	aniversario := r.Group("/aniversario")
	{
		aniversario.GET("/getAniversariosEmpresa", func(c *gin.Context) {
			aniversariantes, err := controller.GetAniversarioEmpresaController()
			if err != nil {
				c.JSON(404, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, aniversariantes)
		})
	}

	r.Run(":8080")
}
