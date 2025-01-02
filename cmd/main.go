package main

import (
	_ "github.com/lib/pq"
	api "gitlab.com/applications2285147/api-go/api/router"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
	"gitlab.com/applications2285147/api-go/database/repository"
	infra "gitlab.com/applications2285147/api-go/infrastructure"
	"gitlab.com/applications2285147/api-go/services"
)

func main() {
	// Criação das dependências
	database := infra.ConstructorConnectDatabase()
	repo := repository.ConstructorConnectDatabase(database)
	aniversarioEmpresaService := services.ConstructorIAniversarioEmpresaRepositorys(repo)
	ctrl := controller.ConstructorIAniversarianteEmpresaServices(aniversarioEmpresaService)
	router := api.ConstructorGetAniversarioEmpresaController(ctrl)

	// Inicia o router
	router.Router(nil)
}
