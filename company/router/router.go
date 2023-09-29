package router

import (
	"company/controller"
	"company/repository"
	"company/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(dbConnectionObjectInit *sqlx.DB) *gin.Engine {
	router := gin.Default()
	registerRoutes(router, dbConnectionObjectInit)
	return router
}

func registerRoutes(registerRoutesRouter *gin.Engine, dbConnectionObjectRegisterRoutes *sqlx.DB) {

	empServiceToRepository := repository.EmployeeRepositoryToRouter(dbConnectionObjectRegisterRoutes)
	empControllerToService := service.EmployeeServiceToRouter(empServiceToRepository)
	empRouterToController := controller.EmployeeControllerToRouter(empControllerToService)

	routeGroup := registerRoutesRouter.Group("/api/v1/")
	{
		routeGroup.GET("/events", empRouterToController.GetEmployees)
	}
}
