package dell

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

type Chassis struct {
	ODataContext     string `json:"@odata.context"`
	ODataID          string `json:"@odata.id"`
	ODataType        string `json:"@odata.type"`
	Actions          ChassisActions
	Assembly         Link
	AssetTag         string `json:"AssetTag"`
	ChassisType      string `json:"ChassisType"`
	Description      string `json:"Description"`
	Id               string `json:"Id"`
	IndicatorLED     string `json:"IndicatorLED"`
	Links            ChassisLinks
	Location         ChassisLocation
	Manufacturer     string `json:"Manufacturer"`
	Model            string `json:"Model"`
	Name             string `json:"Name"`
	NetworkAdapters  Link
	Oem              ChassisOEM
	PartNumber       string `json:"PartNumber"`
	PhysicalSecurity ChassisPhysicalSecurity
	Power            Link
	PowerState       string `json:"PowerState"`
	SKU              string `json:"SKU"`
	SerialNumber     string `json:"SerialNumber"`
	Status           ChassisStatus
	Thermal          Link
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
	ch <- prometheus.MustNewConstMetric(
		ChasState,
		prometheus.GaugeValue,
		chas.StatusToNumber(),
		chas.PartNumber,
		chas.SKU,
		chas.SerialNumber,
	)
	// for i := range sys {
	// 	fmt.Println(sys[i].UUID)
	// }
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

// NewMetrics return a DellHandler struct
func NewMetrics() Chassis {
	return Chassis{}
}
