package main

import (
	_ "github.com/lib/pq"                                                               // Import the PostgreSQL driver for database connectivity.
	api "gitlab.com/applications2285147/api-go/api/router"                              // Import the router package for API routing.
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController" // Import the controller package.
	controllerScreenshot "gitlab.com/applications2285147/api-go/controller/screenshotControllers"
	"gitlab.com/applications2285147/api-go/database/repository"  // Import the repository package for database interaction.
	"gitlab.com/applications2285147/api-go/handlers"             // Import the handlers package for business logic.
	infra "gitlab.com/applications2285147/api-go/infrastructure" // Import the infrastructure package for database connection setup.
	"gitlab.com/applications2285147/api-go/infrastructure/queue"
	"gitlab.com/applications2285147/api-go/services" // Import the services package for business logic.
)

func main() {
	// Create the necessary dependencies for the application.

	// Step 1: Initialize the database connection using the infrastructure package.
	// This sets up the connection to the PostgreSQL database.
	database := infra.ConstructorConnectDatabase()

	// Step 2: Create the repository instance, passing the database connection.
	// This repository is responsible for executing SQL queries and handling database logic.
	repoAniversarioVida := repository.ConstructorAniversariantesVidaConnectionDatabase(database)
	repoAniversarioEmpresa := repository.ConstructorAniversariantesEmpresaConnectionDatabase(database)
	// Step 3: Create the service layer for the aniversarioEmpresa functionality.
	// The service layer provides business logic and integrates with the repository.
	aniversarioVidaService := services.ConstructorAniversariantesVidaRepositorys(repoAniversarioVida)
	aniversarioEmpresaService := services.ConstructorIAniversarioEmpresaRepositorys(repoAniversarioEmpresa)

	// Step 3: Create the services
	screenshotService := &services.ScreenshotService{}

	// Depois passe para o constructor do service
	screenshotQueue := queue.NewScreenshotQueue(screenshotService)
	screenshotServiceInstance := services.ConstructorScreenshotService(screenshotService, screenshotQueue)

	// Step 4: Create the controllers
	ctrlEmpresa := controller.ConstructorIAniversarianteEmpresaServices(aniversarioEmpresaService)
	ctrlVida := controller.ConstructorAniversariantesVidaServices(aniversarioVidaService)
	ctrlScreenshot := controllerScreenshot.ConstructorIScreenshotServices(screenshotServiceInstance)
	// Step 5: Create the handlers
	handlerEmpresa := handlers.ConstructorGetAniversarioEmpresaController(ctrlEmpresa)
	handlerVida := handlers.ConstructorAniversariantesVidaController(ctrlVida)
	handlerScreenshot := handlers.ConstructorScreenshotController(ctrlScreenshot)

	// Step 6: Initialize the handlers
	aniversariosHandler := api.ConstructorAniversariantesHandler(handlerEmpresa, handlerVida)
	screenshotHandler := api.ConstructorScreenshotHandler(handlerScreenshot)

	// Router setup
	routerHandler := api.NewRouterHandler(
		aniversariosHandler,
		screenshotHandler,
	)

	router := routerHandler.SetupRouter()
	router.Run(":8080")

}
