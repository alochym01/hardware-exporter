package base

import "github.com/prometheus/client_golang/prometheus"

var (
	// SysState => System Health Metric
	SysState = prometheus.NewDesc(
		"system_state",
		"system_state {0: OK, 1: Warning, 2: Critical}",
		[]string{"sku", "serialnumber"},
		nil,
	)

	// // ChasPower => Chassis Power Metric
	// ChasPower = prometheus.NewDesc(
	// 	"power_consumed",
	// 	"power_consumed {0: OK, 1: Warning, 2: Critical}",
	// 	[]string{"partnumber", "sku", "serialnumber"},
	// 	nil,
	// )

	// SysStorage => System Storage Metric
	SysStorage = prometheus.NewDesc(
		"storage_status",
		"storage_status {0: OK, 1: Warning, 2: Critical}",
		[]string{"partnumber", "sku", "serialnumber"},
		nil,
	)
)
