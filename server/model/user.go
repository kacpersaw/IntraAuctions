package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kacpersaw/intra-auctions/config"
	"time"
)

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

func (u *User) GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
		"uid":      u.ID,
		"username": u.Username,
		"email":    u.Email,
		"admin":    u.Admin,
	})
	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	return tokenString, err
}
