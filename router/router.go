package router

import (
	"strings"

	"github.com/alochym01/hardware-exporter/domain/server"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/gin-gonic/gin"
)

// Router ...
// func Router(db *sql.DB) *gin.Engine {
func Router(ginMode string) *gin.Engine {
	redfish.ClientDELL = redfish.NewAPIClient("root", "calvin")
	redfish.ClientHPE = redfish.NewAPIClient("username", "password")

	router := gin.Default()

	if strings.Contains(ginMode, "release") {
		gin.SetMode(gin.ReleaseMode)
	}

	dellHandler := server.NewDellHandler()
	router.GET("/metrics/dell", dellHandler.Metric)
	hpeHandler := server.NewHPEHandler()
	router.GET("/metrics/hpe", hpeHandler.Metric)

	return router
}
