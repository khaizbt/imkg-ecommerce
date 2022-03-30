package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/service"
)

type productController struct {
	product_service service.ProductService
}

func ProductController(productService service.ProductService) *productController {
	return &productController{productService}
}

func (h *productController) CreateProduct(c *gin.Context) {
	var input entity.ProductInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}
	fmt.Println("Check 1")
	input.UserID = c.MustGet("loggedUser").(model.User).ID
	err = h.product_service.InsertProduct(input)

	if err != nil {
		responseError := helper.APIResponse("Create Product Failed #CTM61", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Product has been created", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}

func (h *productController) Search(c *gin.Context) {

	//Mapping Input
	input := entity.SearchInput{
		SKU:      c.Query("sku"),
		Title:    c.Query("search"),
		Category: c.Query("category"),
		Brand:    c.Query("brand"),
	}

	result, err := h.product_service.GetProduct(input)

	if err != nil {
		responseError := helper.APIResponse("Search Product Failed #SRC002", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Search Successes", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *productController) UpdateProduct(c *gin.Context) {
	var input entity.ProductInput

	inputID := entity.IdUserInput{}
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		responseError := helper.APIResponse("Update Product Failed #UPD001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	err = c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Update Product Failed #UPD002", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	input.UserID = c.MustGet("currentUser").(model.User).ID //TODO Buat Function Check Jika Prodduct Bukan Miliknya maka g bisa
	err = h.product_service.UpdateProduct(strconv.Itoa(inputID.ID), input)

	if err != nil {
		responseError := helper.APIResponse("Update Product Failed #CTM61", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Product has been updated", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}

func (h *productController) DeleteProduct(c *gin.Context) {
	var input entity.ProductInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Delete Product Failed #DPT001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	err = h.product_service.DeleteProduct(input)

	if err != nil {
		responseError := helper.APIResponse("Delete Product Failed #DPT91", http.StatusBadRequest, "fail", nil)
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	response := helper.APIResponse("Product has been deleted", http.StatusOK, "success", true)
	c.JSON(http.StatusOK, response)
}
