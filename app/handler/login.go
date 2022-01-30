package handler

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/app/auth"
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
	g.POST("/token", loginHandler.Login)
	g.POST("/refresh-token", loginHandler.Login)
}

type LoginHandler interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
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

	var credential auth.LoginCredentials
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
	user, err := h.loginService.LoginUser(credential.Email, credential.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}
	token, refreshToken := h.jWtService.GenerateToken(user, true)
	userResponse := auth.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
		"user":          userResponse,
	})
}

func (h *loginHandler) RefreshToken(c *gin.Context) {

}
