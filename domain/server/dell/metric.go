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
		fmt.Println("=============================")
		fmt.Println("=====SystemsCollection=======")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}
	err = json.Unmarshal(data, &sysCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("=====SystemsCollection=======")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}

	// Set a systems url
	for _, v := range sysCollection.Members {
		redfish.Client.SysURL = fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
	}

	// get System Data
	sysData, err := redfish.Client.Get(redfish.Client.SysURL)
	// Problem connect to server
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("==========Systems============")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}
	var sys Systems
	err = json.Unmarshal(sysData, &sys)
	// Data cannot convert System struct
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("==========Systems============")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}

	// Set system_health metrics
	m.sysHealth(ch, sys)

	// Set all Links
	// Set system storage collection Data
	storeCollectionLink := fmt.Sprintf("%s%s", redfish.Client.Host, sys.Storage.ODataID)
	storeCollectionData, err := redfish.Client.Get(storeCollectionLink)
	// Problem connect to server
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("=Systems storage collection==")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		return
	}
	var storeCollection StorageCollection
	err = json.Unmarshal(storeCollectionData, &storeCollection)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("=Systems storage collection==")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		return
	}

	// get Storage Data
	// Set Storage URL
	var storeURL string
	for _, v := range storeCollection.Members {
		if strings.Contains(v.ODataID, "RAID") {
			storeURL = fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
			fmt.Println(storeURL)
			break
		}
	}
	storeData, err := redfish.Client.Get(storeURL)
	// Problem connect to server
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("====Systems storage=====")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		return
	}

	var store Storage
	err = json.Unmarshal(storeData, &store)
	// Data cannot convert Storage struct
	if err != nil {
		fmt.Println("=============================")
		fmt.Println("====Systems storage=====")
		fmt.Println("=============================")
		fmt.Println(err.Error())
		return
	}
	// Set storage_status metric
	// m.sysStorageStatus(ch, store)

	// Set storage_disk metric
	for _, v := range store.Drives {
		// store.Drives[0].ODataID
		diskURL := fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
		diskData, err := redfish.Client.Get(diskURL)
		// Problem connect to server
		if err != nil {
			fmt.Println("=============================")
			fmt.Println("====Systems storage DISK=====")
			fmt.Println("=============================")
			fmt.Println(err.Error())
			continue
		}
		var disk StorageDisk
		err = json.Unmarshal(diskData, &disk)
		if err != nil {
			fmt.Println("=============================")
			fmt.Println("====Systems storage DISK=====")
			fmt.Println("=============================")
			fmt.Println(err.Error())
			continue
		}
		if disk.PredictedMediaLifeLeftPercent > 0 {
			m.sysStorageDisk(ch, disk)
		}
	}
	// storeVolLink := store.Volumes.ODataID

	// Set system_health metric

	// biosLink := sys.Bios.ODataID
	// ethIfacesLink := sys.EthernetInterfaces.ODataID
	// networkIfacesLink := sys.NetworkInterfaces.ODataID
	// biosLink := sys.Bios.ODataID
	// Set Metrics
}

// Get System Data start
// func (m Metrics) sysData(url string, c redfish.APIClient) (*Systems, error) {
// 	data, err := c.Get(c.SysURL)
// 	// Problem connect to server
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	var sys Systems
// 	err = json.Unmarshal(data, &sys)
// 	// b, _ := json.MarshalIndent(sys, "", "    ")
// 	// fmt.Println(string(b))
// 	// Data cannot convert System struct
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	return &sys, nil
// }

func (m Metrics) sysStorageStatus(ch chan<- prometheus.Metric, s Storage) {
	// d, _ := json.MarshalIndent(s, "", "    ")
	// fmt.Println(string(d))
	ch <- prometheus.MustNewConstMetric(
		base.SysStorageState,
		prometheus.GaugeValue,
		s.StatusToNumber(),
	)
}

func (m Metrics) sysStorageDisk(ch chan<- prometheus.Metric, s StorageDisk) {
	// d, _ := json.MarshalIndent(s, "", "    ")
	// fmt.Println(string(d))
	ch <- prometheus.MustNewConstMetric(
		base.SysStorageDisk,
		prometheus.GaugeValue,
		s.PredictedMediaLifeLeftPercent,
		fmt.Sprintf("%d", s.PhysicalLocation.PartLocation.LocationOrdinalValue),
		fmt.Sprintf("%d", s.CapacityBytes),
		s.Protocol,
		s.MediaType,
	)
}

// Get System Data end

// Set System Metrics start
func (m Metrics) sysHealth(ch chan<- prometheus.Metric, s Systems) {
	// d, _ := json.MarshalIndent(s, "", "    ")
	// fmt.Println(string(d))
	ch <- prometheus.MustNewConstMetric(
		base.SysState,
		prometheus.GaugeValue,
		s.StatusToNumber(),
		s.SKU,
		s.SerialNumber,
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
