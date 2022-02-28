package main

import (
	"math/rand"
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
	[]string{"status_code", "path", "method"},
)

func getStatusCode() int {
	r := rand.Intn(100)
	if r < 50 {
		return 200
	}
	if r < 80 {
		return 404
	}
	if r < 95 {
		return 418
	}
	return 500
}

func main() {
	prometheus.MustRegister(counter_requests)
	prometheus.MustRegister(promInfo)
	promInfo.Set(1)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(404, r.URL.Path)
		counter_requests.WithLabelValues(
			"404",
			r.URL.Path,
			r.Method,
		).Inc()
		w.WriteHeader(404)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		statusCode := getStatusCode()
		fmt.Println(r.Method, statusCode, r.URL.Path)
		counter_requests.WithLabelValues(
			strconv.Itoa(statusCode),
			r.URL.Path,
			r.Method,
		).Inc()
		w.WriteHeader(statusCode)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
