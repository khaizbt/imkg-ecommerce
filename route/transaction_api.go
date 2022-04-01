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
	api.POST("add-cart", middleware.AuthMiddlewareUser(authService, userService, 5), transactionController.AddCart)
	api.GET("list-cart", middleware.AuthMiddlewareUser(authService, userService, 10), transactionController.ListCart)
	api.POST("update-cart", middleware.AuthMiddlewareUser(authService, userService, 6), transactionController.UpdateCart)
	api.POST("delete-cart", middleware.AuthMiddlewareUser(authService, userService, 7), transactionController.DeleteCart)

	//Transaction
	api.POST("checkout", middleware.AuthMiddlewareUser(authService, userService, 8), transactionController.Checkout)
	api.GET("list", middleware.AuthMiddlewareUser(authService, userService, 9), transactionController.ListTransaction)
	api.GET("list-all", middleware.AuthMiddlewareUser(authService, userService, 10), transactionController.ListSalesTransaction) //FIXME Middleware Feature dibuat Khusus Only Penjual
	api.POST("update-status", middleware.AuthMiddlewareUser(authService, userService, 12), transactionController.UpdateStatus)   //FIXME Middleware Feature dibuat Khusus Only Penjual
}
