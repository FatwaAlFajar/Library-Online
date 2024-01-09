package books

import (
	"fmt"
	"net/http"
	"time"

	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Books struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	City      string `json:"city"`
	Year      string `json:"year"`
	Isbn      string `json:"isbn"`
	Quantity  int    `json:"quantity"`
}

// POST /book/tambah
func CreateBook(c *gin.Context) {
	var input Books
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	created_at := time.Now().UTC().Add(7 * time.Hour)

	data := models.Books{
		ID:        id,
		Title:     input.Title,
		Author:    input.Author,
		Publisher: input.Publisher,
		City:      input.City,
		Year:      input.Year,
		Isbn:      input.Isbn,
		Quantity:  input.Quantity,
		CreatedAt: created_at,
		IsDeleted: false,
	}

	fmt.Println(input.Title)

	models.DB.Create(&data)

	c.JSON(http.StatusCreated, gin.H{"data": data})
}

// GET /books/all
func GetAllBooks(c *gin.Context) {
	var book []models.Books
	if err := models.DB.Find(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile records not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
