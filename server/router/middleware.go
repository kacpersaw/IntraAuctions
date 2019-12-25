package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/kacpersaw/intra-auctions/config"
	"github.com/kacpersaw/intra-auctions/model"
	"github.com/rs/cors"
	"net/http"
	"strings"
)

func CommonMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func CorsMiddleware(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Accept", "Content-Length", "X-Requested-With", "Origin"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
	})
	return c.Handler(handler)
}

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var user model.User
		if model.DB.Where("id = ?", token.Claims.(jwt.MapClaims)["uid"].(string)).First(&user).RecordNotFound() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		context.Set(r, "user", user)

		handler.ServeHTTP(w, r)
	})
}
