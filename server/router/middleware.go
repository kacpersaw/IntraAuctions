package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/kacpersaw/intra-auctions/config"
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

		context.Set(r, "uid", token.Claims.(jwt.MapClaims)["uid"].(string))
		context.Set(r, "admin", token.Claims.(jwt.MapClaims)["admin"].(bool))

		handler.ServeHTTP(w, r)
	})
}

func AdminMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		admin := context.Get(r, "admin").(bool)
		if !admin {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
