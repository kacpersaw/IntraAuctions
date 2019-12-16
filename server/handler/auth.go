package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gitlab.com/kacpersaw/intra-auctions/config"
	"gitlab.com/kacpersaw/intra-auctions/ldap"
	"gitlab.com/kacpersaw/intra-auctions/model"
	"gitlab.com/kacpersaw/intra-auctions/util"
	"net/http"
)

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginRequest AuthLoginRequest
	err := decoder.Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		util.MustEncode(json.NewEncoder(w), util.GenericResponse{
			Code:    10001,
			Message: "Invalid body",
		})
		return
	}

	ok, ldapUser, err := ldap.LDAP.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		logrus.Errorf("Error authenticating user %s: %+v", loginRequest.Username, err)

		w.WriteHeader(http.StatusUnauthorized)
		util.MustEncode(json.NewEncoder(w), util.GenericResponse{
			Code:    10002,
			Message: "Invalid username or password",
		})
		return
	}

	if !ok {
		logrus.Errorf("Authenticating failed for user %s", loginRequest.Username)
		w.WriteHeader(http.StatusInternalServerError)
		util.MustEncode(json.NewEncoder(w), util.GenericResponse{
			Code:    10003,
			Message: "Invalid error occurred",
		})
		return
	}

	user := model.User{
		Username: loginRequest.Username,
		Email:    ldapUser[config.LDAP_MAIL_ATTRIBUTE],
		Admin:    false,
	}

	model.DB.Where("username = ?", loginRequest.Username).FirstOrCreate(&user)

	token, err := user.GenerateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.MustEncode(json.NewEncoder(w), util.GenericResponse{
			Code:    10003,
			Message: "Invalid error occurred",
		})
		return
	}

	util.MustEncode(json.NewEncoder(w), TokenResponse{Token: token})
}
