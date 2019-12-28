package model

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Auction struct {
	ID uuid.UUID `gorm:"type:varchar(36);primary_key;column:id" json:"id"`

	Active bool `gorm:"column:active" json:"active"`

	Title            string `gorm:"column:title" json:"title"`
	ShortDescription string `gorm:"column:short_description" json:"short_description"`
	Description      string `gorm:"column:description" json:"description"`

	StartPrice  float32 `gorm:"column:start_price" json:"start_price"`
	MinimalBid  float32 `gorm:"column:minimal_bid" json:"minimal_bid"`
	ActualPrice float32 `gorm:"column:actual_price" json:"actual_price"`

	StartDate time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate   time.Time `gorm:"column:end_date" json:"end_date"`

	Images []Image `gorm:"foreignkey:AuctionID "json:"images"`

	Bids []Bid `gorm:"foreignkey:AuctionID" json:"bids"`

	// Created at timestamp
	CreatedAt time.Time `json:"created_at"`

	// Updated at timestamp
	UpdatedAt time.Time `json:"-"`

	// Deleted at timestamp
	DeletedAt *time.Time `json:"-"`
}

func (a *Auction) BeforeCreate(scope *gorm.Scope) error {
	uuidGen, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}
	scope.SetColumn("ID", uuidGen)

	return nil
}
