// api/router.go
package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handler "gitlab.com/applications2285147/api-go/handler/aniversarioHandlers"
)

func Router(db *sql.DB) {
	r := gin.Default()

	aniversario := r.Group("/aniversario")
	{
		aniversario.GET("/getAniversariosEmpresa", handler.GetAniversariantesHandler)
	}

	r.Run(":8080")
}
