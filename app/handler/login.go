package handler

import (
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"github.com/Fadhli12/go-gin-gorm-playground/service"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigLogin struct {
	R *gin.Engine
}

func NewLoginHandler(c *ConfigLogin, db *gorm.DB) {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginHandler LoginController = LoginHandler(loginService, jwtService)

	g := c.R.Group("/auth")
	g.POST("/login", func(context *gin.Context) {
		token := loginHandler.Login(context)
		if token != "" {
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			context.JSON(http.StatusUnauthorized, nil)
		}
	})
}

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (h *loginController) Login(ctx *gin.Context) string {
	var credential model.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := h.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return h.jWtService.GenerateToken(credential.Email, true)

	}
	return ""
}
