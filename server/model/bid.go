package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Bid struct {
	ID int `gorm:"primary_key" json:"id"`

	Bid float32 `gorm:"bid" json:"bid"`

	Uid string `gorm:"uid" json:"uid"`

	AuctionID uuid.UUID `gorm:"type:varchar(36);column:auction_id;" json:"auction_id"`

	// Created at timestamp
	CreatedAt time.Time `json:"created_at"`

	// Updated at timestamp
	UpdatedAt time.Time `json:"-"`

	// Deleted at timestamp
	DeletedAt *time.Time `json:"-"`
}
