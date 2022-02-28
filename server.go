package main

import (
	"net/http"

	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server started.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
