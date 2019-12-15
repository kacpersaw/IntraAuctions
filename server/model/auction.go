package model

import "time"

type Auction struct {
	ID int `gorm:"primary_key" json:"id"`

	Active bool `gorm:"active" json:"active"`

	Title            string `gorm:"title" json:"title"`
	ShortDescription string `gorm:"short_description" json:"short_description"`
	Description      string `gorm:"description" json:"description"`

	StartPrice  float32 `gorm:"start_price" json:"start_price"`
	MinimumBid  float32 `gorm:"minimum_bid" json:"minimum_bid"`
	ActualPrice float32 `gorm:"actual_price" json:"actual_price"`

	StartDate time.Time `gorm:"start_date" json:"start_date"`
	EndDate   time.Time `gorm:"end_date" json:"end_date"`

	// Created at timestamp
	CreatedAt time.Time `json:"created_at"`

	// Updated at timestamp
	UpdatedAt time.Time `json:"-"`

	// Deleted at timestamp
	DeletedAt *time.Time `json:"-"`
}
