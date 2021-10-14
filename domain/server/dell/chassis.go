package dell

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alochym01/hardware-exporter/domain/server/base"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

type Chassis struct {
	base.Meta
	base.Chassis
	base.Actions
	Assembly         base.Link
	Description      string `json:"Description"`
	Links            ChassisLinks
	Location         ChassisLocation
	Oem              ChassisOEM
	PCIeSlots        base.Link
	PartNumber       string `json:"PartNumber"`
	PhysicalSecurity ChassisPhysicalSecurity
	Sensors          base.Link
	Status           ChassisStatus
	UUID             string `json:"UUID"`
}

// Describe a description of metrics
func (s Chassis) Describe(ch chan<- *prometheus.Desc) {
	ch <- ChasState
}

// Collect return a metric with all desc value and metric value
func (s Chassis) Collect(ch chan<- prometheus.Metric) {
	// var sysOK float64 = 0.0
	var chasOK float64 = 0.0
	// sysOK = 0
	// chasOK = 0
	var chasCollection ChassisCollection
	var chas Chassis
	dataCollection, err := redfish.Client.Get()

	// Problem connect to server
	if err != nil {
		// sysOK = 2
		chasOK = 2
		fmt.Println(err.Error())
		// ch <- prometheus.MustNewConstMetric(
		// 	config.SysState,
		// 	prometheus.GaugeValue,
		// 	sysOK,
		// )
		ch <- prometheus.MustNewConstMetric(
			ChasState,
			prometheus.GaugeValue,
			chasOK,
			"",
			"",
			"",
		)
		return
	}
	// m, err := goFish.Get(config.Config.Endpoint)
	// sys, syserr := m.Systems()
	// chas, cherr := m.Chassis()

	// for _, s := range sys {
	// 	fmt.Printf("%#v", s)
	// }
	// for _, c := range chas {
	// 	fmt.Printf("%#v", c)
	// }
	// if syserr != nil || cherr != nil {
	// 	if syserr != nil {
	// 		sysOK = 2
	// 		fmt.Println(err.Error())
	// 	}
	// 	if cherr != nil {
	// 		chasOK = 2
	// 		fmt.Println(err.Error())
	// 	}
	// 	ch <- prometheus.MustNewConstMetric(
	// 		config.SysState,
	// 		prometheus.GaugeValue,
	// 		sysOK,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		config.ChasState,
	// 		prometheus.GaugeValue,
	// 		chasOK,
	// 	)
	// 	return
	// }
	// ch <- prometheus.MustNewConstMetric(
	// 	config.SysState,
	// 	prometheus.GaugeValue,
	// 	sysOK,
	// )
	// err = json.NewDecoder(strings.NewReader(data)).Decode(&chas)
	err = json.Unmarshal(dataCollection, &chasCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			ChasState,
			prometheus.GaugeValue,
			chasOK,
			"",
			"",
			"",
		)
		return
	}
	for _, v := range chasCollection.Members {
		if strings.Contains(v.ODataID, "System") {
			// Set a chassis url
			redfish.Client.URL = fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
			data, err := redfish.Client.Get()
			// Problem connect to server
			if err != nil {
				// sysOK = 2
				chasOK = 2
				fmt.Println(err.Error())
				// ch <- prometheus.MustNewConstMetric(
				// 	config.SysState,
				// 	prometheus.GaugeValue,
				// 	sysOK,
				// )
				ch <- prometheus.MustNewConstMetric(
					ChasState,
					prometheus.GaugeValue,
					chasOK,
					"",
					"",
					"",
				)
				return
			}
			err = json.Unmarshal(data, &chas)
			// Data cannot convert Chassis struct
			if err != nil {
				ch <- prometheus.MustNewConstMetric(
					ChasState,
					prometheus.GaugeValue,
					chasOK,
					"",
					"",
					"",
				)
				return
			}
		}
	}
	// Everything is ok
	ch <- prometheus.MustNewConstMetric(
		ChasState,
		prometheus.GaugeValue,
		chas.StatusToNumber(),
		chas.PartNumber,
		chas.SKU,
		chas.SerialNumber,
	)
}

func (c Chassis) StatusToNumber() float64 {
	switch c.Status.Health {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	default:
		return 3.0
	}
}

// NewMetrics return a Chassis struct
func NewMetrics() Chassis {
	return Chassis{}
}
