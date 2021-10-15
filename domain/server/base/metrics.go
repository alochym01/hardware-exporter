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
	SysStorageState = prometheus.NewDesc(
		"storage_state",
		"storage_state {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)

	SysStorageDisk = prometheus.NewDesc(
		"storage_drive_ssd_endurance",
		"storage_drive_ssd_endurance {0: OK, 1: Warning, 2: Critical}",
		[]string{"id", "capacity", "interface_type", "media_type"},
		nil,
	)
)
