package borrowing

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Pinjaman struct {
	ID_Anggota string `json:"nik"`
	ID_Book    string `json:"isbn"`
}

// POST /borrow/tambah
func CreateBorrow(c *gin.Context) {
	var input Pinjaman
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	created_at := time.Now().UTC().Add(7 * time.Hour)

	data := models.Borrowing{
		ID:              id,
		ReturningStatus: false,
		ID_Anggota:      input.ID_Anggota,
		ID_Book:         input.ID_Book,
		CreatedAt:       created_at,
		IsDeleted:       false,
	}

	models.DB.Create(&data)

	c.JSON(http.StatusCreated, gin.H{"data": data})
}

// GET /borrow/all
func GetAllBorrowing(c *gin.Context) {
	var data []models.Borrowing
	if err := models.DB.Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile records not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
