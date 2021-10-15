package router

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/alochym01/hardware-exporter/domain/server"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/gin-gonic/gin"
)

// Router ...
// func Router(db *sql.DB) *gin.Engine {
func Router(ginMode string) *gin.Engine {
	// Create a custom Transport
	// The default value of Transport's MaxIdleConnsPerHost.
	// const DefaultMaxIdleConnsPerHost = 2
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.IdleConnTimeout = 60 * time.Second
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	// Disable SSL check
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// Init APIClient config
	redfish.Client = redfish.NewAPIClient(
		&http.Client{
			Transport: transport,
			Timeout:   time.Duration(10) * time.Second,
		},
	)

	router := gin.Default()

	gin.SetMode(ginMode)

	dellHandler := server.NewDellHandler()
	router.GET("/metrics", dellHandler.Metric)

	return router
}
