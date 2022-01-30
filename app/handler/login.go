package handler

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/app/auth"
	"github.com/Fadhli12/go-gin-gorm-playground/common"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLoginHandler(R *gin.Engine, db *gorm.DB) {
	var loginService auth.LoginService = auth.NewLoginService(db)
	var jwtService auth.JWTService = auth.JWTAuthService()
	var loginHandler LoginHandler = authLoginHandler(loginService, jwtService)

	g := R.Group("/auth")
	g.POST("/", loginHandler.Login)
}

type LoginHandler interface {
	Login(ctx *gin.Context)
}

type loginHandler struct {
	loginService auth.LoginService
	jWtService   auth.JWTService
}

func authLoginHandler(loginService auth.LoginService,
	jWtService auth.JWTService) LoginHandler {
	return &loginHandler{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (h *loginHandler) Login(c *gin.Context) {

	var credential model.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	isUserAuthenticated := h.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		token := h.jWtService.GenerateToken(credential.Email, true)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"errors": common.ErrorRequest("user password not correct", http.StatusUnauthorized),
	})
}
