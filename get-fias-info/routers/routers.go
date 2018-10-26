package routers

import (
	"encoding/json"
	"fmt"
	"go-games/get-fias-info/fias"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	fiasURL      = "https://fias.nalog.ru/WebServices/Public/DownloadService.asmx"
	fiasEnvelope = "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" " +
		"xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">  " +
		"<soap12:Body><GetAllDownloadFileInfo xmlns=\"http://fias.nalog.ru/WebServices/Public/DownloadService.asmx\" />" + "</soap12:Body></soap12:Envelope>"
)

// StartWebServer starts web server with port
func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	router := NewRouter()
	err := http.ListenAndServe(":"+port, router) // Goroutine will block here
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

// Route Defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{
	Route{
		"GetFiasInfo", // Name
		"GET",         // HTTP method
		"/fias/info",  // Route pattern
		func(w http.ResponseWriter, r *http.Request) {
			result, err := fias.GetFias(fiasURL, fiasEnvelope)
			if err != nil {
				fmt.Println(err)
			}

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			js, err := json.Marshal(result)
			if err != nil {
				fmt.Println("json error")
				return
			}

			fmt.Println("result object is: " + string(js))

			w.Write([]byte(js))
		},
	},
}

//NewRouter returns new router
func NewRouter() *mux.Router {
	// Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)
	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, route := range routes {
		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
