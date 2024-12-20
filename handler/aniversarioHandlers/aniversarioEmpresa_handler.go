// handler/aniversario_handler.go
package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/applications2285147/api-go/controller"
	"gitlab.com/applications2285147/api-go/database"
)

func GetAniversariantesHandler(c *gin.Context) {

	//O NewAniversarioController() é um padrão de construção em Go chamado "constructor function" (função construtora).
	//Cria uma nova instância da struct AniversarioController
	//Inicializa ela com as dependências necessárias (neste caso, a conexão com o banco)
	//Retorna um ponteiro para essa nova instância
	aniversarioController := controller.ConstructorAniversarioController(database.DB)

	aniversariantes, err := aniversarioController.BuscarAniversariantesDoDia()

	if err != nil {

		c.JSON(404, gin.H{"error": "Nenhum valor encontrado" + err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"aniversariantes": aniversariantes,
	})

}
