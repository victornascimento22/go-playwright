package main

import (
	_ "github.com/lib/pq"                                                               // Import the PostgreSQL driver for database connectivity.
	api "gitlab.com/applications2285147/api-go/api/router"                              // Import the router package for API routing.
	controller "gitlab.com/applications2285147/api-go/controller/aniversarioController" // Import the controller package.
	"gitlab.com/applications2285147/api-go/database/repository"                         // Import the repository package for database interaction.
	infra "gitlab.com/applications2285147/api-go/infrastructure"                        // Import the infrastructure package for database connection setup.
	"gitlab.com/applications2285147/api-go/services"                                    // Import the services package for business logic.
)

func main() {
	// Create the necessary dependencies for the application.

	// Step 1: Initialize the database connection using the infrastructure package.
	// This sets up the connection to the PostgreSQL database.
	database := infra.ConstructorConnectDatabase()

	// Step 2: Create the repository instance, passing the database connection.
	// This repository is responsible for executing SQL queries and handling database logic.
	repo := repository.ConstructorConnectDatabase(database)

	// Step 3: Create the service layer for the aniversarioEmpresa functionality.
	// The service layer provides business logic and integrates with the repository.
	aniversarioEmpresaService := services.ConstructorIAniversarioEmpresaRepositorys(repo)

	// Step 4: Create the controller layer.
	// The controller manages the interaction between the service layer and the HTTP router.
	ctrl := controller.ConstructorIAniversarianteEmpresaServices(aniversarioEmpresaService)

	// Step 5: Initialize the router with the controller.
	// The router sets up HTTP routes and handles incoming API requests.
	router := api.ConstructorGetAniversarioEmpresaController(ctrl)

	// Start the HTTP server and listen for incoming requests.
	// The nil argument indicates no explicit database instance is passed to the router.
	router.Router(nil)
}
