package server

import (
	"fmt"

	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// DellHandler ...
type DellHandler struct{}

// Metric ...
func (handler DellHandler) Metric(c *gin.Context) {
	url := fmt.Sprintf("%s%s", c.Query("host"), "/redfish/v1")

	// Set URL get from Request
	redfish.Client.URL = url
	data, _ := redfish.Client.Get()
	fmt.Println(string(data))

	// Make promhttp response to Request
	h := promhttp.Handler()
	h.ServeHTTP(c.Writer, c.Request)
}

// NewDellHandler return a DellHandler struct
func NewDellHandler() *DellHandler {
	return &DellHandler{}
}
