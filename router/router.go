package router

import (
	"github.com/alochym01/hardware-exporter/domain/server"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/gin-gonic/gin"
)

// Router ...
// func Router(db *sql.DB) *gin.Engine {
func Router(ginMode string) *gin.Engine {
	redfish.Client = redfish.NewAPIClient()

	router := gin.Default()

	gin.SetMode(ginMode)

	dellHandler := server.NewDellHandler()
	router.GET("/metrics", dellHandler.Metric)

	return router
}
