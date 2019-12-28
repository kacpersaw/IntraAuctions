package router

import (
	"github.com/gorilla/mux"
	"github.com/kacpersaw/intra-auctions/config"
	"github.com/kacpersaw/intra-auctions/handler"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.PathPrefix("/images/").
		Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(config.IMG_DIR))))

	v1 := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		var h http.Handler
		h = CommonMiddleware(route.HandlerFunc)

		if route.AuthRequired {
			h = AuthMiddleware(h)
		}

		v1.Path(route.Pattern).Handler(h).Name(route.Name).Methods(route.Method)
	}

	return router
}

type Route struct {
	Name          string
	Method        string
	Pattern       string
	HandlerFunc   http.HandlerFunc
	AuthRequired  bool
	AdminRequired bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
		false,
		false,
	},

	//Auth
	Route{
		"Login",
		"POST",
		"/auth/login",
		handler.Login,
		false,
		false,
	},

	//Auctions
	Route{
		"Create auction",
		"POST",
		"/auction",
		handler.AuctionCreate,
		true,
		true,
	},
	Route{
		"List auctions",
		"GET",
		"/auction",
		handler.AuctionList,
		true,
		true,
	},
	Route{
		"Get active auction",
		"GET",
		"/auction/active",
		handler.AuctionGetActive,
		true,
		false,
	},
	Route{
		"Delete auction",
		"DELETE",
		"/auction/{id}",
		handler.AuctionDelete,
		true,
		true,
	},
	Route{
		"Set active auction",
		"PUT",
		"/auction/{id}/active",
		handler.AuctionSetActive,
		true,
		true,
	},
}
