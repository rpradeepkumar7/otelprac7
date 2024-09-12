package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"

    //"otelprac7/logging"
    //"otelprac7/metrics"
    //"otelprac7/tracing"

    "github.com/rpradeepkumar7/otelprac7/logging"
    "github.com/rpradeepkumar7/otelprac7/metrics"
    "github.com/rpradeepkumar7/otelprac7/tracing"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Set up tracing
    tracerProvider, err := tracing.InitTracer()
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err := tracerProvider.Shutdown(context.Background()); err != nil {
            log.Fatal(err)
        }
    }()

    // Set up logging
    hostname, ipAddress, macAddress, osType, osVersion := logging.GetSystemInfo()
    firewallStatus := logging.GetFirewallStatus()
    networkLatency := logging.GetNetworkLatency()

    logEntry := logging.NewLogEntry(hostname, ipAddress, macAddress, osType, osVersion, firewallStatus, networkLatency)
    logEntryJSON, err := json.MarshalIndent(logEntry, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Log Entry in JSON format:")
    log.Println(string(logEntryJSON))

    // Set up Prometheus metrics
    registry := metrics.InitMetrics()
    http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
    go func() {
        log.Println("Serving metrics at :2112/metrics")
        log.Fatal(http.ListenAndServe(":2112", nil))
    }()

    log.Println("OpenTelemetry is set up and running!")
}
