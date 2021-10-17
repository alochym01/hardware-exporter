package server

import (
	"github.com/alochym01/hardware-exporter/domain/server/hpe"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// DellHandler ...
type DellHandler struct{}

// Metric ...
func (handler DellHandler) Metric(c *gin.Context) {
	// Set Host get from Request
	redfish.Client.Server = c.Query("host")

	// Register Server Dell Metrics
	// Using custom registry
	registry := prometheus.NewRegistry()
	// dellMetrics := dell.NewMetrics()
	server := hpe.NewMetrics()
	registry.MustRegister(server)

	// Make promhttp response to Request
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(c.Writer, c.Request)
}

// NewDellHandler return a DellHandler struct
func NewDellHandler() *DellHandler {
	return &DellHandler{}
}
