package handler

import (
	"encoding/json"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/kacpersaw/intra-auctions/config"
	"github.com/kacpersaw/intra-auctions/events"
	"github.com/kacpersaw/intra-auctions/model"
	"github.com/kacpersaw/intra-auctions/util"
	"github.com/segmentio/ksuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type AuctionRequest struct {
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`

	StartPrice float32 `json:"start_price"`
	MinimalBid float32 `json:"minimal_bid"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type BidRequest struct {
	Bid float32 `json:"bid"`
}

func AuctionCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(100000)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := AuctionRequest{}
	err = json.Unmarshal([]byte(r.FormValue("auction")), &data)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auction := model.Auction{
		Active:           false,
		Title:            data.Title,
		ShortDescription: data.ShortDescription,
		Description:      data.Description,
		StartPrice:       data.StartPrice,
		MinimalBid:       data.MinimalBid,
		ActualPrice:      data.StartPrice,
		StartDate:        data.StartDate,
		EndDate:          data.EndDate,
	}
	model.DB.Save(&auction)

	m := r.MultipartForm
	files := m.File["images"]
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		filename := ksuid.New().String() + filepath.Ext(files[i].Filename)

		destination, err := os.Create(config.IMG_DIR + "/" + filename)
		defer destination.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(destination, file); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		image := model.Image{
			Url:       "/images/" + filename,
			AuctionID: auction.ID,
		}
		model.DB.Save(&image)
	}

	w.WriteHeader(http.StatusOK)
	util.MustEncode(json.NewEncoder(w), auction)
}

func AuctionDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	model.DB.Where("id = ?", params["id"]).Delete(model.Auction{})
	w.WriteHeader(http.StatusNoContent)
}

func AuctionList(w http.ResponseWriter, r *http.Request) {
	var auctions []model.Auction
	model.DB.Order("id desc").Find(&auctions)

	w.WriteHeader(http.StatusOK)
	util.MustEncode(json.NewEncoder(w), auctions)
}

func AuctionSetActive(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	model.DB.Model(&model.Auction{}).Where("active = ?", true).Update("active", false)
	model.DB.Model(&model.Auction{}).Where("id = ?", params["id"]).Update("active", true)

	w.WriteHeader(http.StatusNoContent)
}

func AuctionGetActive(w http.ResponseWriter, r *http.Request) {
	var auction model.Auction
	model.DB.Where("active = ?", true).First(&auction)
	model.DB.Model(&auction).Related(&auction.Images)

	w.WriteHeader(http.StatusOK)
	util.MustEncode(json.NewEncoder(w), auction)
}

func AuctionDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var auction model.Auction
	model.DB.Where("id = ?", params["id"]).First(&auction)
	model.DB.Model(&auction).Related(&auction.Images)
	model.DB.Model(&auction).Related(&auction.Bids)

	w.WriteHeader(http.StatusOK)
	util.MustEncode(json.NewEncoder(w), auction)
}

func AuctionBid(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var auction model.Auction
	if model.DB.Where("id = ?", params["id"]).First(&auction).RecordNotFound() {
		logrus.Error("Could not find given auction")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()
	if currentTime.Before(auction.StartDate) || currentTime.After(auction.EndDate) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := BidRequest{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data.Bid <= 0 || data.Bid < (auction.ActualPrice+auction.MinimalBid) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bid := model.Bid{
		Bid:       data.Bid,
		Uid:       context.Get(r, "uid").(string),
		AuctionID: auction.ID,
	}
	model.DB.Save(&bid)

	auction.ActualPrice = data.Bid
	model.DB.Save(&auction)

	auctionStr, _ := json.Marshal(auction)
	events.SSE.SendMessage("/v1/events/"+auction.ID.String(), sse.SimpleMessage(string(auctionStr)))

	w.WriteHeader(http.StatusOK)
}
