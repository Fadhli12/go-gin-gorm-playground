package handler

import (
	"errors"
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/app/user"
	"github.com/Fadhli12/go-gin-gorm-playground/common"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(R *gin.Engine, db *gorm.DB) {

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	h := &userHandler{userService}

	g := R.Group("/user")
	g.GET("/", h.GetUsers)
	g.GET("/:id", h.GetUser)
	g.POST("/", h.CreateUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}

//func NewUserHandler(userService user.Service) *userHandler {
//	return &userHandler{userService}
//}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	var userResponses []user.UserResponse
	for _, item := range users {
		userResponse := convertToUserResponse(item)
		userResponses = append(userResponses, userResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": userResponses,
	})
}

func (h *userHandler) GetUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	userById, err := h.userService.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"errors": common.ErrorRequest("not found", http.StatusNotFound),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	userResponse := convertToUserResponse(userById)
	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var userPost user.UserPost
	err := c.ShouldBind(&userPost)
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
	createdUser, err := h.userService.Create(userPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	userResponse := convertToUserResponse(createdUser)
	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var userUpdate user.UserUpdate
	err := c.ShouldBind(&userUpdate)
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
	user, err := h.userService.Update(id, userUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	userResponse := convertToUserResponse(user)
	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	user, err := h.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	userResponse := convertToUserResponse(user)
	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func convertToUserResponse(b model.User) user.UserResponse {
	return user.UserResponse{
		ID:    b.ID,
		Name:  b.Name,
		Email: b.Email,
	}
}
