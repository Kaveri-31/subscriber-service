package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Total API Requests
var SubscriberRequestsTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "subscriber_requests_total",
		Help: "Total subscriber API requests",
	},
)

// Total Subscribers Created
var SubscriberCreateTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "subscriber_create_total",
		Help: "Total subscribers created",
	},
)

// Total Subscribers Updated
var SubscriberUpdateTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "subscriber_update_total",
		Help: "Total subscribers updated",
	},
)

// Total Subscribers Deleted
var SubscriberDeleteTotal = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "subscriber_delete_total",
		Help: "Total subscribers deleted",
	},
)

func RegisterMetrics() {
	prometheus.MustRegister(
		SubscriberRequestsTotal,
		SubscriberCreateTotal,
		SubscriberUpdateTotal,
		SubscriberDeleteTotal,
	)
}
