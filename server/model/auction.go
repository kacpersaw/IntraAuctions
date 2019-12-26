package model

import "time"

type Auction struct {
	ID int `gorm:"primary_key" json:"id"`

	Active bool `gorm:"column:active" json:"active"`

	Title            string `gorm:"column:title" json:"title"`
	ShortDescription string `gorm:"column:short_description" json:"short_description"`
	Description      string `gorm:"column:description" json:"description"`

	StartPrice  float32 `gorm:"column:start_price" json:"start_price"`
	MinimumBid  float32 `gorm:"column:minimum_bid" json:"minimum_bid"`
	ActualPrice float32 `gorm:"column:actual_price" json:"actual_price"`

	StartDate time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate   time.Time `gorm:"column:end_date" json:"end_date"`

	Images []Image `gorm:"foreignkey:AuctionID"json:"images"`

	// Created at timestamp
	CreatedAt time.Time `json:"created_at"`

	// Updated at timestamp
	UpdatedAt time.Time `json:"-"`

	// Deleted at timestamp
	DeletedAt *time.Time `json:"-"`
}
