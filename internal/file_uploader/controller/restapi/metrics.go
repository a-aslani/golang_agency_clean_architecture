package restapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (r *controller) RegisterMetrics(serviceName string) {

	r.Router.GET("/metrics", prometheusHandler())

	r.reqCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "http_request_counter",
		Name:      serviceName,
		Help:      fmt.Sprintf("Count of request to the %s service", serviceName),
	})

	r.reqLatency = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "http_request_latency",
		Name:      serviceName,
		Buckets:   []float64{0.1, 0.5, 1.0},
	})
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
