package handler

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/kacpersaw/intra-auctions/model"
	"github.com/kacpersaw/intra-auctions/util"
	"net/http"
	"time"
)

type AuctionRequest struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`

	StartPrice float32 `json:"start_price"`
	MinimumBid float32 `json:"minimum_bid"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func AuctionCreate(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(model.User)
	if !user.Admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data := AuctionRequest{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auction := model.Auction{
		Active:           false,
		Title:            data.Title,
		ShortDescription: data.ShortDescription,
		Description:      data.Description,
		StartPrice:       data.StartPrice,
		MinimumBid:       data.MinimumBid,
		ActualPrice:      0,
		StartDate:        data.StartDate,
		EndDate:          data.EndDate,
	}
	model.DB.Save(auction)

	w.WriteHeader(http.StatusOK)
	util.MustEncode(json.NewEncoder(w), auction)
}

func AuctionDelete(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(model.User)
	if !user.Admin {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	model.DB.Where("id = ?", params["id"]).Delete(model.Auction{})
	w.WriteHeader(http.StatusNoContent)
}
