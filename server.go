package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var promInfo = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "example_info",
	Help: "Instance info",
	ConstLabels: prometheus.Labels{
		"version": "0.1.0-dev",
		"branch":  "master",
	},
})

var counter_requests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "example",
		Name:      "requests_total",
		Help:      "Total nuber of requests",
	})

func main() {
	prometheus.MustRegister(counter_requests)
	prometheus.MustRegister(promInfo)
	promInfo.Set(1)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		counter_requests.Inc()
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
