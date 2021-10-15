package dell

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alochym01/hardware-exporter/domain/server/base"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
}

func (m Metrics) Desc(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
}

// Describe a description of metrics
func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
}

// Collect return a metric with all desc value and metric value
func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	// System Metrics
	m.SystemsCollector(ch, *redfish.Client)
	// Chassis Metrics
	// m.ChassisCollector(ch, *redfish.Client)
}

// NewMetrics return a Metrics struct
func NewMetrics() Metrics {
	return Metrics{}
}

func (m Metrics) SystemsCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// var sysCollection
	var sysCollection SystemsCollection
	url := redfish.Client.SysURL
	data, err := redfish.Client.Get(url)
	// Problem connect to server
	if err != nil {
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(
			base.SysState,
			prometheus.GaugeValue,
			2.0,
			"",
			"",
		)
		return
	}
	err = json.Unmarshal(data, &sysCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(
			base.SysState,
			prometheus.GaugeValue,
			2.0,
			"",
			"",
		)
		return
	}

	// Set a systems url
	for _, v := range sysCollection.Members {
		redfish.Client.SysURL = fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
	}
	// get System Data
	sysData, err := m.sysData(redfish.Client.SysURL, *redfish.Client)
	if err != nil {
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(
			base.SysState,
			prometheus.GaugeValue,
			2.0,
			"",
			"",
		)
		return
	}
	b, _ := json.MarshalIndent(sysData, "", "    ")
	fmt.Println(string(b))

	// Set system_health metrics
	m.sysHealth(ch, *sysData)

	// Set all Links
	storageLink := sysData.Storage.ODataID
	// biosLink := sysData.Bios.ODataID
	// ethIfacesLink := sysData.EthernetInterfaces.ODataID
	// networkIfacesLink := sysData.NetworkInterfaces.ODataID
	// biosLink := sysData.Bios.ODataID
	// Set Metrics
}

// Get System Data start
func (m Metrics) sysData(url string, c redfish.APIClient) (*Systems, error) {
	data, err := c.Get(c.SysURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var sys Systems
	err = json.Unmarshal(data, &sys)
	// b, _ := json.MarshalIndent(sys, "", "    ")
	// fmt.Println(string(b))
	// Data cannot convert System struct
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &sys, nil
}
func (m Metrics) sysStorage(url string, c redfish.APIClient) (*Systems, error) {
	data, err := c.Get(c.SysURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var sys Systems
	err = json.Unmarshal(data, &sys)
	// b, _ := json.MarshalIndent(sys, "", "    ")
	// fmt.Println(string(b))
	// Data cannot convert System struct
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &sys, nil
}

// Get System Data end

// Set System Metrics start
func (m Metrics) sysHealth(ch chan<- prometheus.Metric, sys Systems) {
	ch <- prometheus.MustNewConstMetric(
		base.SysState,
		prometheus.GaugeValue,
		sys.StatusToNumber(),
		sys.SKU,
		sys.SerialNumber,
	)
}

// Set System Metrics end

// sysHealth metrics from Systems
// func (m Metrics) sysHealth(ch chan<- prometheus.Metric, c redfish.APIClient) {
// 	data, err := c.Get(c.SysURL)
// 	// Problem connect to server
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		ch <- prometheus.MustNewConstMetric(
// 			base.SysState,
// 			prometheus.GaugeValue,
// 			2.0,
// 			"",
// 			"",
// 			"",
// 		)
// 		return
// 	}
// 	var sys Systems
// 	err = json.Unmarshal(data, &sys)
// 	// b, _ := json.MarshalIndent(sys, "", "    ")
// 	// fmt.Println(string(b))
// 	// Data cannot convert ChassisCollection struct
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		ch <- prometheus.MustNewConstMetric(
// 			base.SysState,
// 			prometheus.GaugeValue,
// 			2.0,
// 			"",
// 			"",
// 			"",
// 		)
// 		return
// 	}
// 	// Everything is ok
// 	ch <- prometheus.MustNewConstMetric(
// 		base.SysState,
// 		prometheus.GaugeValue,
// 		sys.StatusToNumber(),
// 		sys.PartNumber,
// 		sys.SKU,
// 		sys.SerialNumber,
// 	)
// }

func (m Metrics) ChassisCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// var chasCollection
	var chasCollection ChassisCollection
	url := redfish.Client.ChasURL
	data, err := redfish.Client.Get(url)
	// Problem connect to server
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(data, &chasCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Set a chassis url
	for _, v := range chasCollection.Members {
		if strings.Contains(v.ODataID, "System") {
			redfish.Client.ChasURL = fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
			fmt.Println(redfish.Client.ChasURL)
			break
		}
	}

	// get Chassis Data
	dataChassis, err := c.Get(c.ChasURL)
	if err != nil {
		return
	}

	// Data cannot convert Chassis struct
	var chas Chassis
	err = json.Unmarshal(dataChassis, &chas)
	if err != nil {
		return
	}
	b, _ := json.MarshalIndent(chas, "", "    ")
	fmt.Println(string(b))
	// m.sysHealth(ch, *redfish.Client)
}

func (m Metrics) chasPowerControl(ch chan<- prometheus.Metric, c redfish.APIClient) {
	data, err := c.Get(c.ChasURL)
	// Problem connect to server
	if err != nil {
		// fmt.Println(err.Error())
		// ch <- prometheus.MustNewConstMetric(
		// 	base.ChasPower,
		// 	prometheus.GaugeValue,
		// 	2.0,
		// 	"",
		// 	"",
		// 	"",
		// )
		return
	}
	var chas Chassis
	err = json.Unmarshal(data, &chas)
	// b, _ := json.MarshalIndent(sys, "", "    ")
	// fmt.Println(string(b))
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(
			base.SysState,
			prometheus.GaugeValue,
			2.0,
			"",
			"",
			"",
		)
		return
	}
	// Everything is ok
	// ch <- prometheus.MustNewConstMetric(
	// 	base.ChasPower,
	// 	prometheus.GaugeValue,
	// 	sys.StatusToNumber(),
	// 	sys.PartNumber,
	// 	sys.SKU,
	// 	sys.SerialNumber,
	// )
}
