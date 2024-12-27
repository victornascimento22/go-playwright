// api/router.go
package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

type IAniversariantesEmpresaController struct {
	controller controller.IAniversarioEmpresaController
}

func ConstructorGetAniversarioEmpresaController(ctrl controller.IAniversarioEmpresaController) *IAniversariantesEmpresaController {
	return &IAniversariantesEmpresaController{
		controller: ctrl,
	}
}

func (x *IAniversariantesEmpresaController) Router(db *sql.DB) {
	r := gin.Default()

	aniversario := r.Group("/aniversario")
	{
		aniversario.GET("/getAniversariosEmpresa", func(c *gin.Context) {
			aniversariantes, err := x.controller.GetAniversarioEmpresaController()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, aniversariantes)
		})
	}

	r.Run(":8080")
}
