package handler

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/author"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type authorHandler struct {
	authorService author.Service
}

type ConfigAuthor struct {
	R *gin.Engine
}

func NewAuthorHandler(c *ConfigAuthor, db *gorm.DB) {

	authorRepository := author.NewRepository(db)
	authorService := author.NewService(authorRepository)
	h := &authorHandler{authorService}

	g := c.R.Group("/author")
	g.GET("/", h.GetAuthors)
	g.GET("/:id", h.GetAuthor)
	g.POST("/", h.CreateAuthor)
	g.PUT("/:id", h.UpdateAuthor)
	g.DELETE("/:id", h.DeleteAuthor)
}

//func NewAuthorHandler(authorService author.Service) *authorHandler {
//	return &authorHandler{authorService}
//}

func (h *authorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.authorService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	var authorResponses []author.AuthorResponse
	for _, item := range authors {
		authorResponse := convertToAuthorResponse(item)
		authorResponses = append(authorResponses, authorResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": authorResponses,
	})
}

func (h *authorHandler) GetAuthor(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	authorById, err := h.authorService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	authorResponse := convertToAuthorResponse(authorById)
	c.JSON(http.StatusOK, gin.H{
		"data": authorResponse,
	})
}

func (h *authorHandler) CreateAuthor(c *gin.Context) {
	var authorPost author.AuthorPost
	err := c.Bind(&authorPost)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errorMessages,
		})
		return
	}
	author, err := h.authorService.Create(authorPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": author,
	})
}

func (h *authorHandler) UpdateAuthor(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var authorUpdate author.AuthorUpdate
	err := c.ShouldBindJSON(&authorUpdate)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errorMessages,
		})
	}
	author, err := h.authorService.Update(id, authorUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	authorResponse := convertToAuthorResponse(author)
	c.JSON(http.StatusOK, gin.H{
		"data": authorResponse,
	})
}

func (h *authorHandler) DeleteAuthor(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	author, err := h.authorService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	authorResponse := convertToAuthorResponse(author)
	c.JSON(http.StatusOK, gin.H{
		"data": authorResponse,
	})
}

func convertToAuthorResponse(b author.Author) author.AuthorResponse {
	return author.AuthorResponse{
		ID:        b.ID,
		Name:      b.Name,
		Email:     b.Email,
		Biography: b.Biography,
	}
}
