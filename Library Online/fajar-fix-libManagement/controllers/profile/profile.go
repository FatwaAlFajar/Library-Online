package profile

import (
	"net/http"

	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type IniInput struct {
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

// POST /anggota/tambah
func CreateMember(c *gin.Context) {
	var input IniInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()

	// created_at := time.Now().UTC().Add(7 * time.Hour)

	data := models.Anggota{
		ID:      id,
		Nik:     input.Nik,
		Name:    input.Name,
		Address: input.Address,
		Email:   input.Email,
	}

	models.DB.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GET /profile/all
func GetAllProfile(c *gin.Context) {
	var profiles []models.Anggota
	if err := models.DB.Find(&profiles).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile records not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profiles})
}

// // GET /profile/:id
// func GetProfile(c *gin.Context) {
// 	// Fetch user based on the user ID from the route parameter
// 	var user models.Users
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User record not found"})
// 		return
// 	}

// 	// Fetch profile with eager loading of the associated user
// 	var profile models.Profile
// 	if err := models.DB.Preload("Users").Where("id_user = ?", user.ID).First(&profile).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile record not found for the user"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		}
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": profile})
// }

// // PUT /profile/ubah/:id
// func UpdateProfile(c *gin.Context) {
// 	var profile models.Profile
// 	if err := models.DB.Where("id_user = ?", c.Param("id")).First(&profile).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Profile record not found for the user"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		}
// 		return
// 	}

// 	var input IniInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	data := models.Profile{
// 		Name:    input.Name,
// 		Nik:     input.Nik,
// 		Address: input.Address,
// 	}

// 	models.DB.Model(&profile).Updates(data)

// 	c.JSON(http.StatusOK, gin.H{"data": profile})
// }
