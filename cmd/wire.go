//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	api "gitlab.com/applications2285147/api-go/api/router"
	controller "gitlab.com/applications2285147/api-go/controller"
	"gitlab.com/applications2285147/api-go/database/repository"
	"gitlab.com/applications2285147/api-go/handlers"
	infra "gitlab.com/applications2285147/api-go/infrastructure"
	"gitlab.com/applications2285147/api-go/services"
)

type Handler struct {
	AniversarioEmpresaHandler handlers.IAniversariantesEmpresaHandler
	AniversarioVidaHandler    handlers.IAniversariantesVidaHandler
	ScreenshotHandler         handlers.IScreenshotHandler
	WebsocketHandler          *handlers.WebsocketHandlerImpl
	ScreenshotService         services.IScreenshotService
	RouterHandler             *api.RouterHandler
}

type Controller struct {
	AniversarioEmpresaController controller.IAniversarioEmpresaController
	AniversarioVidaController    controller.IAniversariantesVidaController
	ScreenshotController         controller.IScreenshotController
}

type Repositories struct {
	RepositoriesEmpresa repository.IAniversariantesEmpresaRepository
	RepositoriesVida    repository.IAniversariantesVidaRepository
}

type Services struct {
	ServicesEmpresa    services.IAniversarioEmpresaServices
	ServicesVida       services.IAniversariantesVidaServices
	ServicesScreenshot services.IScreenshotService
}

type Router struct {
	Router api.IRouter
}

func ProvideDatabase() infra.IConnectDatabase {
	return infra.ConstructorConnectDatabase()
}

func ProvideRouter(handlers2 Handler) (*gin.Engine, error) {
	aniversariantesHandler := &api.AniversariantesHandler{
		EmpresaHandler: handlers2.AniversarioEmpresaHandler,
		VidaHandler:    handlers2.AniversarioVidaHandler,
	}

	screenshotHandler := handlers2.ScreenshotHandler

	websocketHandler := &api.IWS{
		WebsocketHandler: handlers2.WebsocketHandler,
	}

	routerHandler := api.NewRouterHandler(aniversariantesHandler, screenshotHandler, websocketHandler, handlers2.ScreenshotService)

	router, err := routerHandler.SetupRouter()
	if err != nil {
		return nil, err
	}
	return router, nil
}

func ProvideHandlers(ctrl Controller) Handler {
	return Handler{
		AniversarioEmpresaHandler: handlers.ConstructorGetAniversarioEmpresaController(ctrl.AniversarioEmpresaController),
		AniversarioVidaHandler:    handlers.ConstructorAniversariantesVidaController(ctrl.AniversarioVidaController),
		ScreenshotHandler:         handlers.ConstructorIScreenshotController(ctrl.ScreenshotController),
		WebsocketHandler:          handlers.NewWebsocketHandler(&services.ScreenshotService{}),
		ScreenshotService:         &services.ScreenshotService{},
	}
}

func ProvideControllers(serv Services) Controller {
	return Controller{
		AniversarioEmpresaController: controller.ConstructorIAniversarianteEmpresaServices(serv.ServicesEmpresa),
		AniversarioVidaController:    controller.ConstructorAniversariantesVidaServices(serv.ServicesVida),
		ScreenshotController:         controller.ConstructorIScreenshotServices(serv.ServicesScreenshot),
	}
}

func ProvideRepositories(database infra.IConnectDatabase) Repositories {
	return Repositories{
		RepositoriesEmpresa: repository.ConstructorAniversariantesEmpresaConnectionDatabase(database),
		RepositoriesVida:    repository.ConstructorAniversariantesVidaConnectionDatabase(database),
	}
}

func ProvideServices(repos Repositories) Services {
	return Services{
		ServicesEmpresa: services.ConstructorIAniversarioEmpresaRepositorys(repos.RepositoriesEmpresa),
		ServicesVida:    services.ConstructorAniversariantesVidaRepositorys(repos.RepositoriesVida),
	}
}

func InitializeApplication() (*gin.Engine, error) {
	wire.Build(
		ProvideDatabase,
		ProvideRepositories,
		ProvideServices,
		ProvideControllers,
		ProvideHandlers,
		ProvideRouter,
	)
	return nil, nil // Este retorno é ignorado pelo Wire
}
