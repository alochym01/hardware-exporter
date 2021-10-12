package dell

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// SysState => System Health Metric
	// SysState = prometheus.NewDesc(
	// 	"system_state",
	// 	"system_state {0: OK, 1: Warning, 2: Critical}",
	// 	[]string{},
	// 	nil,
	// )

	// ChasState => System Health Metric
	ChasState = prometheus.NewDesc(
		"chassis_state",
		"chassis_state {0: OK, 1: Warning, 2: Critical}",
		[]string{"partnumber", "sku", "serialnumber"},
		nil,
	)
)
