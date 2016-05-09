package currypoint

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/newtechfellas/CurryPoint/handlerfuncs"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"home",
		"GET",
		"/",
		handlerfuncs.Home,
	},
	Route{
		"login",
		"POST",
		"/login",
		handlerfuncs.Login,
	},
	Route{
		"NewUser",
		"POST",
		"/NewUser",
		handlerfuncs.NewUser,
	},
	Route{
		"UpdateUser",
		"POST",
		"/UpdateUser",
		handlerfuncs.UpdateUser,
	},
	Route{
		"PlaceOrder",
		"POST",
		"/order",
		handlerfuncs.PlaceOrder,
	},
	Route{
		"FetchOrder",
		"GET",
		"/order/{id}",
		handlerfuncs.GetOrder,
	},
}