package handler

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type bookHandler struct {
	bookService book.Service
}

type Config struct {
	R *gin.Engine
}

func NewBookHandler(c *Config, db *gorm.DB) {

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
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
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

		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errorMessages,
		})
		return
	}
	book, err := h.bookService.Create(bookPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	var bookUpdate book.BookUpdate
	err := c.ShouldBindJSON(&bookUpdate)
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
	book, err := h.bookService.Update(id, bookUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	bookResponse := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	bookResponse := convertToBookResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
