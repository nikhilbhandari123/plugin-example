package routers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nikhilbhandari123/plugin-example/internal/mypackage"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	myRouter := router.PathPrefix("/api/test").Subrouter()
	for _, route := range vpnroutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		myRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

// Add route related to VPN portal only.
var vpnroutes = Routes{
	Route{
		"HttpCaller",
		strings.ToUpper("Post"),
		"/testfunction",
		mypackage.HttpCaller,
	},
}
