package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var
	REQUEST_COUNT = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_app_requests_count",
		Help: "Total App HTTP Requests count.",
	})

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			name := vars["name"]
			greetings := fmt.Sprintf("Happy Birthday %s :)", name)
			rw.Write([]byte(greetings))

			REQUEST_COUNT.Inc()
	}).Methods("GET")

	log.Println("Starting the application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":8000", router)
}