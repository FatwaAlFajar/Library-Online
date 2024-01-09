package models

import (
	"time"
)

type Books struct {
	ID        string `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	City      string `json:"city"`
	Year      string `json:"year"`
	Isbn      string `json:"isbn"`
	Quantity  int    `json:"quantity"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Borrowing struct {
	ID              string `json:"id" gorm:"primary_key"`
	ReturningStatus bool   `json:"returning_status"`

	// foreign keys
	ID_Anggota string  `json:"id_anggota"`
	Anggota    Anggota `json:"anggota" gorm:"foreignKey:ID_Anggota"`
	ID_Book    string  `json:"id_book"`
	Books      Books   `json:"books" gorm:"foreignKey:ID_Book"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}
