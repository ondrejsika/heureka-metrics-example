package main

import (
	"net/http"
	"strconv"

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

var counter_requests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "example",
		Name:      "requests_total",
		Help:      "Total nuber of requests",
	},
	[]string{"status_code"},
)

func main() {
	prometheus.MustRegister(counter_requests)
	prometheus.MustRegister(promInfo)
	promInfo.Set(1)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode := 200
		fmt.Println(r.URL.Path)
		counter_requests.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		w.WriteHeader(statusCode)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
