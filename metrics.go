package main

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var filterTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "filter_total",
		Help: "filtering node requests total counter",
	},
	[]string{"code", "api"},
)

func init() {
	prometheus.MustRegister(filterTotal)
}

func addMetricsHandlers(mux *http.ServeMux) {
	mux.Handle("/metrics", promhttp.Handler())
}

func incFilterTotal(statusCode int, api string) {
	filterTotal.WithLabelValues(strconv.Itoa(statusCode), api).Inc()
}
