package route

import (
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/controller"
	"github.com/khaizbt/imkg-ecommerce/middleware"
	"github.com/khaizbt/imkg-ecommerce/repository"
	"github.com/khaizbt/imkg-ecommerce/service"
)

func ProductRoute(route *gin.Engine, userService service.UserService) {
	productRepo := repository.NewElasticRepo()
	productService := service.NewProductService(productRepo, repository.NewRepository())
	productHandler := controller.ProductController(productService)

	api := route.Group("/api/v1/product")
	api.GET("search", productHandler.Search)
	api.GET("detail/:product_id", productHandler.DetailProduct)
	api.POST("create", middleware.AuthMiddlewareUser(authService, userService, 1), productHandler.CreateProduct)
	api.POST("update/:id", middleware.AuthMiddlewareUser(authService, userService, 2), productHandler.UpdateProduct)
	api.POST("delete", middleware.AuthMiddlewareUser(authService, userService, 3), productHandler.DeleteProduct)
}
