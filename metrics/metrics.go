// metrics.go
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Client-side metrics:
	//
	// 1. Number of Requests Sent: This is the total number of requests that
	// the client has made. This will help understand the total load
	// the client is trying to distribute.
	// 
	// 2. Request Distribution: Record the server (pod) each request is sent
	// to. This helps evaluate how evenly requests are distributed across the
	// servers.
	// 
	// 3. Request Latencies: This measures the time taken to complete each
	// request. High latencies might indicate issues with load balancing if certain
	// servers are becoming overloaded.
	// 
	// 4. Error Rates: The number of requests that result in errors. If some servers
	// consistently return more errors than others, this might indicate an imbalance
	// in load distribution.
	// 
	// 5. Time taken to detect changes: Measure how long it takes for the
	// client to detect and adapt to scaling operations (pod
	// addition/removal). This is not a standard metric that can be pulled
	// from a library but will likely need to be designed.
	
	ClientRequestsSent = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "poc_client_requests_sent_total",
		Help: "The total number of requests sent by the client",
	}, []string{"server_id"})

	ClientRequestDistribution = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "poc_client_request_distribution_total",
		Help: "The distribution of requests across servers",
	}, []string{"server_id"})

	ClientRequestLatencies = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "poc_client_request_latency_seconds",
		Help:    "The latency of client requests",
		Buckets: prometheus.DefBuckets,
	}, []string{"server_id"})

	ClientErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "poc_client_errors_total",
		Help: "The total number of errors encountered by the client",
	})

	ClientChangeDetectionTimes = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "poc_client_change_detection_time_seconds",
		Help:    "The time it takes for the client to detect and adapt to scaling operations",
		Buckets: prometheus.DefBuckets,
	}, []string{"server_id"})

	ClientInactivityTimes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "poc_client_inactivity_times",
		Help: "The time since the last response from each known server",
	}, []string{"server_id"})

	// Server-side metrics:
	//
	// 1. Number of Requests Received: This measures the number of requests each
	// server (pod) receives, helping understand the load distribution from
	// another perspective.
	// 
	// 2. Response Times: This is the time it takes for the server to process and
	// respond to each request. It can be useful to identify overloaded servers.
	// 
	// 3 .Error Rates: This is the number of errors occurring on the
	// server-side. This could indicate a server is overloaded or experiencing other
	// issues.
	// 
	// 4. CPU/Memory Utilization: This helps to determine if any server is being
	// overloaded with requests.
	
	ServerRequestsReceived = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "poc_server_requests_received_total",
		Help: "The total number of requests received by the server",
	}, []string{"server_id"})

	ServerResponseTimes = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "poc_server_response_time_seconds",
		Help:    "The time it takes for the server to process and respond to each request",
		Buckets: prometheus.DefBuckets,
	}, []string{"server_id"})

	ServerErrors = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "poc_server_errors_total",
		Help: "The total number of errors encountered by the server",
	}, []string{"server_id"})
)
