package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

func InitMetrics() *prometheus.Registry {
    registry := prometheus.NewRegistry()

    httpRequestsTotal := prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint"},
    )

    cpuUsage := prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "cpu_usage",
        Help: "Current CPU usage",
    })

    registry.MustRegister(httpRequestsTotal, cpuUsage)
    return registry
}
