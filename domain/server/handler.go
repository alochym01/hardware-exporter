package server

import (
	"github.com/alochym01/hardware-exporter/domain/server/dell"
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
	redfish.Client.Host = c.Query("host")

	// sysURL := fmt.Sprintf("%s%s", redfish.Client.Host, "/redfish/v1/Systems")
	// chasURL := fmt.Sprintf("%s%s", redfish.Client.Host, "/redfish/v1/Chassis")
	// url := fmt.Sprintf("%s%s", redfish.Client.Host, "/redfish/v1/Chassis")
	// url := fmt.Sprintf("%s%s", c.Query("host"), "/redfish/v1/Chassis")

	// Set URL get from Request
	// redfish.Client.SysURL = sysURL
	// redfish.Client.ChasURL = chasURL

	// fmt.Println("Handler SYS URL -- ", redfish.Client.SysURL)
	// fmt.Println("Handler CHASS URL -- ", redfish.Client.ChasURL)

	// Register Server Dell Metrics
	// Using custom registry
	registry := prometheus.NewRegistry()
	server := dell.NewMetrics()
	registry.MustRegister(server)

	// Make promhttp response to Request
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(c.Writer, c.Request)
}

// NewDellHandler return a DellHandler struct
func NewDellHandler() *DellHandler {
	return &DellHandler{}
}
