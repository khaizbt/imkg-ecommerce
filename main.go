package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/khaizbt/imkg-ecommerce/config"
	"github.com/khaizbt/imkg-ecommerce/middleware"
	"github.com/khaizbt/imkg-ecommerce/repository"
	"github.com/khaizbt/imkg-ecommerce/route"
	"github.com/khaizbt/imkg-ecommerce/service"
)

func main() {
	repo := repository.NewRepository()
	userService := service.NewUserService(repo)

	router := gin.Default()
	router.Use(middleware.SecureMiddleware())
	route.RouteUser(router, userService)
	route.ProductRoute(router, userService)
	route.TransactionRoute(router, userService)
	_ = router.Run(":8000")

}
