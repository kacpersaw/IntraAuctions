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
		var handler http.Handler
		handler = route.HandlerFunc

		v1.Path(route.Pattern).Handler(handler).Name(route.Name).Methods(route.Method)
	}

	return router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},

	Route{
		"Login",
		"POST",
		"/auth/login",
		handler.Login,
	},
}
