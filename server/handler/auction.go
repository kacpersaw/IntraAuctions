package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kacpersaw/intra-auctions/config"
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
	MinimumBid float32 `json:"minimum_bid"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func AuctionCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(100000)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m := r.MultipartForm
	var images []model.Image
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

		images = append(images, model.Image{
			Url: "/images/" + filename,
		})
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
		MinimumBid:       data.MinimumBid,
		ActualPrice:      0,
		StartDate:        data.StartDate,
		EndDate:          data.EndDate,
		Images:           images,
	}
	model.DB.Save(&auction)

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
