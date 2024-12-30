package main

import (
	_ "github.com/lib/pq"
	api "gitlab.com/applications2285147/api-go/api/router"
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController"
	"gitlab.com/applications2285147/api-go/database/repository"
	infra "gitlab.com/applications2285147/api-go/infrastructure"
)

func main() {
	// Criação das dependências
	database := infra.ConstructorConnectDatabase()
	repo := repository.ConstructorConnectDatabase(database)
	ctrl := controller.ConstructorIAniversarianteEmpresaRepository(repo)
	router := api.ConstructorGetAniversarioEmpresaController(ctrl)

	// Inicia o router
	router.Router(nil)
}
