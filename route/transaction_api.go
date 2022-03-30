package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/controller"
	"github.com/khaizbt/imkg-ecommerce/middleware"
	"github.com/khaizbt/imkg-ecommerce/repository"
	"github.com/khaizbt/imkg-ecommerce/service"
)

func TransactionRoute(route *gin.Engine, userService service.UserService) {
	//transactionRepo := repository.NewMongoRepository()
	transactionService := service.NewTransactionService(repository.NewElasticRepo(), repository.NewRepository(), repository.NewMongoRepository())
	transactionController := controller.TransactionController(transactionService)

	api := route.Group("/api/v1/transaction")
	api.POST("add-cart", middleware.AuthMiddlewareUser(authService, userService, 1), transactionController.AddCart)
	api.GET("list-cart", middleware.AuthMiddlewareUser(authService, userService, 1), transactionController.ListCart)
	api.POST("update-cart", middleware.AuthMiddlewareUser(authService, userService, 1), transactionController.UpdateCart)
	api.POST("delete-cart", middleware.AuthMiddlewareUser(authService, userService, 1), transactionController.DeleteCart)
}
