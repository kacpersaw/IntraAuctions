package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kacpersaw/intra-auctions/config"
	"github.com/kacpersaw/intra-auctions/ldap"
	"github.com/kacpersaw/intra-auctions/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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

	groups, err := ldap.LDAP.GetGroupsOfUser(loginRequest.Username)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 48).Unix(),
		"uid":   ldapUser[config.LDAP_USERNAME_ATTRIBUTE],
		"admin": util.StringInSlice(config.LDAP_ADMIN_GROUP_NAME, groups),
	})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		util.MustEncode(json.NewEncoder(w), util.GenericResponse{
			Code:    10003,
			Message: "Invalid error occurred",
		})
		return
	}

	util.MustEncode(json.NewEncoder(w), TokenResponse{Token: tokenString})
}
