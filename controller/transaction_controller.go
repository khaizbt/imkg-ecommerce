package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/service"
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

	input.UserID = c.MustGet("loggedUser").(model.User).ID
	err = h.transaction_service.AddToCart(input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM61", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Product has been created", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) ListCart(c *gin.Context) {
	userId := c.MustGet("loggedUser").(model.User).ID

	results, err := h.transaction_service.ListCart(userId)

	if err != nil {
		responseError := helper.APIResponse("Get Cart Failed #GTM917", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Get Data Cart Success", http.StatusOK, "success", results)
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) UpdateCart(c *gin.Context) {

	var input entity.CartInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	userId := c.MustGet("loggedUser").(model.User).ID

	err = h.transaction_service.UpdateCart(input, userId)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Get Data Cart Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)

}

func (h *transactionController) DeleteCart(c *gin.Context) {
	var input entity.CartInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.UserID = c.MustGet("loggedUser").(model.User).ID

	err = h.transaction_service.DeleteCart(input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Delete Data Cart Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)

}
