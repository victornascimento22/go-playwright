// handler/aniversario_handler.go
package handler

import (
	"github.com/gin-gonic/gin"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
)

func GetAniversariantesEmpresaHandler(c *gin.Context) {

	aniversariantes, err := controller.GetAniversarioEmpresaController()
	if err != nil {

		c.JSON(404, gin.H{"error": "Nenhum valor encontrado" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"aniversariantes": aniversariantes,
	})

}
