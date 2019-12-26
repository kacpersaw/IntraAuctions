package model

import "time"

type Image struct {
	ID int `gorm:"primary_key" json:"id"`

	Url string `gorm:"url" json:"url"`

	AuctionID int `gorm:"column:auction_id" json:"auction_id"`

	// Created at timestamp
	CreatedAt time.Time `json:"created_at"`

	// Updated at timestamp
	UpdatedAt time.Time `json:"-"`

	// Deleted at timestamp
	DeletedAt *time.Time `json:"-"`
}
