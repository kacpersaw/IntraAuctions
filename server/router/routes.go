package router

import (
	"github.com/gorilla/mux"
	"github.com/kacpersaw/intra-auctions/handler"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	v1 := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		var h http.Handler
		h = route.HandlerFunc

		if route.AuthRequired {
			h = AuthMiddleware(h)
		}

		v1.Path(route.Pattern).Handler(h).Name(route.Name).Methods(route.Method)
	}

	return router
}

type Route struct {
	Name         string
	Method       string
	Pattern      string
	HandlerFunc  http.HandlerFunc
	AuthRequired bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
		false,
	},

	Route{
		"Login",
		"POST",
		"/auth/login",
		handler.Login,
		false,
	},

	Route{
		"Create auction",
		"POST",
		"/auction",
		handler.AuctionCreate,
		true,
	},
}
