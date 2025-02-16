package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    RequestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests by method and endpoint",
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    prometheus.MustRegister(RequestCounter)
}