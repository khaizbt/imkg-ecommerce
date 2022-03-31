package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khaizbt/imkg-ecommerce/config"
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/service"
)

type (
	userController struct {
		userService service.UserService
		authService config.AuthService
	}

	UserFormatter struct {
		UserID   int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}
)

func NewUserController(userService service.UserService, authService config.AuthService) *userController {
	return &userController{userService, authService}
}
func FormatUser(user model.User, token string) UserFormatter {
	formatter := UserFormatter{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return formatter
}

func (h *userController) Login(c *gin.Context) {
	var input entity.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		responseError := helper.APIResponse("Login Failed #LOG001", http.StatusUnprocessableEntity, "fail", err.Error())
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	loginUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		responseError := helper.APIResponse("Login Failed #LOG002", http.StatusUnprocessableEntity, "fail", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, responseError)
		return
	}

	token, err := h.authService.GenerateTokenUser(loginUser.ID)

	if err != nil {
		responsError := helper.APIResponse("Login Failed", http.StatusBadGateway, "fail", "Unable to generate token")
		c.JSON(http.StatusBadGateway, responsError)
		return
	}

	//For Set Header if login/logout

	response := helper.APIResponse("Login Success", http.StatusOK, "success", FormatUser(loginUser, token))

	c.JSON(http.StatusOK, response)
}
