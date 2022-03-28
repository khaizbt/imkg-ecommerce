package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/config"
	"github.com/khaizbt/imkg-ecommerce/controller"
	"github.com/khaizbt/imkg-ecommerce/service"
)

var authService = config.NewServiceAuth()

func RouteUser(route *gin.Engine, service service.UserService) {
	userController := controller.NewUserController(service, authService)

	api := route.Group("/api/v1")
	api.POST("/login", userController.Login)
}
