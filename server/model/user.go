package model

import "time"

type User struct {
	ID int `gorm:"primary_key" json:"id"`

	Username string `gorm:"column:username" json:"username"`

	Email string `gorm:"column:email" json:"email"`

	Admin bool `gorm:"column:admin" json:"admin"`

	// Created at timestamp
	CreatedAt time.Time `json:"created_at"`

	// Updated at timestamp
	UpdatedAt time.Time `json:"-"`

	// Deleted at timestamp
	DeletedAt *time.Time `json:"-"`
}
