package models

import "time"

type Users struct {
	ID       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Anggota struct {
	ID      string `json:"id" gorm:"primary_key"`
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}
