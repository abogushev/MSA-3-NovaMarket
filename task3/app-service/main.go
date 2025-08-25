package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "No of request handled",
	},
)

var arr = ""

func main() {
	mx := sync.Mutex{}
	prometheus.MustRegister(reqCounter)

	router := http.NewServeMux()

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqCounter.Inc()
		podName := os.Getenv("HOSTNAME")
		if podName == "" {
			podName = "Неизвестен"
		}
		mx.Lock()
		defer mx.Unlock()
		arr += "AAAAAAAAAA"
		fmt.Fprintf(w, "Идентификатор пода: %s\n", podName)
	})

	http.ListenAndServe(":8080", router)

}
