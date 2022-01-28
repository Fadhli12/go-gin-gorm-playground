package handler

import (
	"errors"
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/app/book"
	"github.com/Fadhli12/go-gin-gorm-playground/common"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type bookHandler struct {
	bookService book.Service
}

type ConfigBook struct {
	R *gin.Engine
}

func NewBookHandler(c *ConfigBook, db *gorm.DB) {

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	h := &bookHandler{bookService}

	g := c.R.Group("/book")
	g.GET("/", h.GetBooks)
	g.GET("/:id", h.GetBook)
	g.POST("/", h.CreateBook)
	g.PUT("/:id", h.UpdateBook)
	g.DELETE("/:id", h.DeleteBook)
}

//func NewBookHandler(bookService book.Service) *bookHandler {
//	return &bookHandler{bookService}
//}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	var bookResponses []book.BookResponse
	for _, item := range books {
		bookResponse := convertToBookResponse(item)
		bookResponses = append(bookResponses, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponses,
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	bookById, err := h.bookService.FindByID(id)
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
		return
	}
	bookResponse := convertToBookResponse(bookById)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookPost book.BookPost
	err := c.Bind(&bookPost)
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
	createdBook, err := h.bookService.Create(bookPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	bookResponse := convertToBookResponse(createdBook)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var bookUpdate book.BookUpdate
	err := c.Bind(&bookUpdate)
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
	updatedBook, err := h.bookService.Update(id, bookUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	bookResponse := convertToBookResponse(updatedBook)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	deletedBook, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	bookResponse := convertToBookResponse(deletedBook)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b model.Book) book.BookResponse {
	author := book.AuthorResponse{
		Name:      b.Author.Name,
		Email:     b.Author.Email,
		Biography: b.Author.Biography,
	}
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
		AuthorID:    b.AuthorID,
		Author:      author,
	}
}
