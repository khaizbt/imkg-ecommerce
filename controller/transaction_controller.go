package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/service"
	"net/http"
)

type transactionController struct {
	transaction_service service.TransactionService
}

func TransactionController(transactionService service.TransactionService) *transactionController {
	return &transactionController{transactionService}
}
func (h *transactionController) AddCart(c *gin.Context) {
	var input entity.CartInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}
	fmt.Println("Check 1")
	input.UserID = c.MustGet("loggedUser").(model.User).ID
	err = h.transaction_service.AddToCart(input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM61", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Product has been created", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}
