// api/router.go
package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"gitlab.com/applications2285147/api-go/handler"
)

func Router(db *sql.DB) {
	r := gin.Default()

	aniversario := r.Group("/aniversario")
	{
		aniversario.GET("/getAniversarios", handler.GetAniversariantesHandler)
	}

	r.Run(":8080")
}
