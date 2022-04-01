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
		responseError := helper.APIResponse("Add To Cart Failed #CRT001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.UserID = c.MustGet("loggedUser").(model.User).ID
	err = h.transaction_service.AddToCart(input)

	if err != nil {
		responseError := helper.APIResponse("Add To Cart Failed  #CRT89", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Cart Has been Added", http.StatusOK, "success", true)
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
		responseError := helper.APIResponse("Update Cart Failed #UPC001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	userId := c.MustGet("loggedUser").(model.User).ID

	err = h.transaction_service.UpdateCart(input, userId)

	if err != nil {
		responseError := helper.APIResponse("Update Cart Failed #UPC076", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Update Data Cart Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)

}

func (h *transactionController) DeleteCart(c *gin.Context) {
	var input entity.CartInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Delete Cart Failed #CTM001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.UserID = c.MustGet("loggedUser").(model.User).ID

	err = h.transaction_service.DeleteCart(input)

	if err != nil {
		responseError := helper.APIResponse("Delete Cart Failed #CTM001", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Delete Data Cart Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)

}

func (h *transactionController) Checkout(c *gin.Context) {
	var input entity.Checkout

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Checkout Failed #CTO001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.UserID = c.MustGet("loggedUser").(model.User).ID

	err = h.transaction_service.Checkout(input)

	if err != nil {
		responseError := helper.APIResponse("Checkout Failed #CTM001", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Checkout Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}

func (h *transactionController) ListTransaction(c *gin.Context) {
	userId := c.MustGet("loggedUser").(model.User).ID

	listTransaction, err := h.transaction_service.ListTransaction(userId)

	if err != nil {
		responseError := helper.APIResponse("Get List Transaction Failed #CTM001", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Get List Transaction Success", http.StatusOK, "success", listTransaction)
	c.JSON(http.StatusOK, response)

}

func (h *transactionController) ListSalesTransaction(c *gin.Context) {

	listTransaction, err := h.transaction_service.ListSalesTransaction()

	if err != nil {
		responseError := helper.APIResponse("Get List Sales Transaction Failed #CTM001", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Get List Sales Transaction Success", http.StatusOK, "success", listTransaction)
	c.JSON(http.StatusOK, response)

}

func (h *transactionController) UpdateStatus(c *gin.Context) {
	var input entity.TransactionStatus

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Update Status Failed #CTO001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	if input.Status != "Dikonfirmasi" && input.Status != "Dikirim" && input.Status != "Dibatalkan" {
		responseError := helper.APIResponse("Update Status Failed #CTO001", http.StatusUnprocessableEntity, "fail", "Format Status Tidak Valid")
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	if input.Status == "Dikirim" && input.Courier.CourierName == "" {
		responseError := helper.APIResponse("Update Status Failed #CTO002", http.StatusUnprocessableEntity, "fail", "Status Dikirim Wajib Melampirkan Data Pengiriman")
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	err = h.transaction_service.UpdateStatus(input)

	if err != nil {
		responseError := helper.APIResponse("Update Status Failed #CPC081", http.StatusBadRequest, "fail", err.Error())
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Update Status Success", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}
