package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var
	REQUEST_INPROGRESS = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "go_app_requests_inprogress",
		Help: "Number of application requests in progress",
	})

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/birthday/{name}", func(rw http.ResponseWriter, r *http.Request) {
			REQUEST_INPROGRESS.Inc()
			vars := mux.Vars(r)
			name := vars["name"]
			greetings := fmt.Sprintf("Happy Birthday %s :)", name)
			time.Sleep(5 * time.Second)
			rw.Write([]byte(greetings))

			REQUEST_INPROGRESS.Dec()
	}).Methods("GET")

	log.Println("Starting the application server...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe(":8000", router)
}