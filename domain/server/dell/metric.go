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
	m.ChassisCollector(ch, *redfish.Client)
}

// NewMetrics return a Metrics struct
func NewMetrics() Metrics {
	return Metrics{}
}

func (m Metrics) SystemsCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// Create a copy redfish.Client
	c = *redfish.Client
	fmt.Println(c.Host)
	// SystemsCollection start
	var sysCollection SystemsCollection
	sysCollectionURL := fmt.Sprintf("%s%s", c.Host, "/redfish/v1/Systems")
	data, err := redfish.Client.Get(sysCollectionURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(sysCollectionURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}
	err = json.Unmarshal(data, &sysCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(sysCollectionURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}
	// SystemsCollection end

	// Systems start
	// Set a systems url
	var sysURL string
	for _, v := range sysCollection.Members {
		sysURL = fmt.Sprintf("%s%s", c.Host, v.ODataID)
	}

	// get System Data
	sysData, err := c.Get(sysURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(sysURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}
	var sys Systems
	err = json.Unmarshal(sysData, &sys)
	// Data cannot convert System struct
	if err != nil {
		fmt.Println(sysURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}

	// Set system_health metrics
	m.sysHealth(ch, sys) // Systems end

	// Set all Links

	// Systems Storage start
	// Systems Storage Collection Data
	storeCollectionLink := fmt.Sprintf("%s%s", c.Host, sys.Storage.ODataID)
	storeCollectionData, err := c.Get(storeCollectionLink)
	// Problem connect to server
	if err != nil {
		fmt.Println(storeCollectionLink)
		fmt.Println(err.Error())
		return
	}
	var storeCollection StorageCollection
	err = json.Unmarshal(storeCollectionData, &storeCollection)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println(storeCollectionLink)
		fmt.Println(err.Error())
		return
	}
	// Systems Storage Collection End

	// Systems Storage Data start
	// Set Systems Storage URL
	var storeURL string
	for _, v := range storeCollection.Members {
		if strings.Contains(v.ODataID, "RAID") {
			// storeURL = fmt.Sprintf("%s%s", redfish.Client.Host, v.ODataID)
			storeURL = fmt.Sprintf("%s%s", c.Host, v.ODataID)
			fmt.Println(storeURL)
			break
		}
	}
	// Get Systems Storage Data
	storeData, err := c.Get(storeURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(storeURL)
		fmt.Println(err.Error())
		return
	}

	var store Storage
	err = json.Unmarshal(storeData, &store)
	// Data cannot convert Storage struct
	if err != nil {
		fmt.Println(storeURL)
		fmt.Println(err.Error())
		return
	}
	// Set storage_status metric
	m.sysStorageStatus(ch, store) // Systems Storage Data end

	// Systems Storage Disk start
	// Set Systems Storage Disk metric
	for _, v := range store.Drives {
		// TODO go routine start
		diskURL := fmt.Sprintf("%s%s", c.Host, v.ODataID)
		// Get Systems Storage Disk Data
		diskData, err := c.Get(diskURL)
		// Problem connect to server
		if err != nil {
			fmt.Println(diskURL)
			fmt.Println(err.Error())
			continue
		}
		var disk StorageDisk
		err = json.Unmarshal(diskData, &disk)
		// Data cannot convert StorageDisk struct
		if err != nil {
			fmt.Println(diskURL)
			fmt.Println(err.Error())
			continue
		}
		// TODO go routine end
		// Check Disk is SSD
		if disk.PredictedMediaLifeLeftPercent > 0 {
			m.sysStorageDisk(ch, disk)
		}
	}
	// Systems Storage Disk end

	// Systems Ethernet Interfaces start
	// Systems Ethernet Interfaces Collection start
	// Set Systems Ethernet Interfaces Collection URL
	ethIfacesLink := fmt.Sprintf("%s%s", c.Host, sys.EthernetInterfaces.ODataID)
	// Systems Ethernet Interfaces Collection
	etherData, err := c.Get(ethIfacesLink)
	if err != nil {
		fmt.Println(ethIfacesLink)
		fmt.Println(err.Error())
		return
	}
	var ethIf EthernetInterfaceCollection
	err = json.Unmarshal(etherData, &ethIf)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println(ethIfacesLink)
		fmt.Println(err.Error())
		return
	}
	// Systems Ethernet Interfaces Collection end
	// Systems Ethernet Interfaces start
	for _, v := range ethIf.Members {
		// TODO go routine start
		ifURL := fmt.Sprintf("%s%s", c.Host, v.ODataID)
		// Get Ethernet Interfaces Data
		ifData, err := c.Get(ifURL)
		// Problem connect to server
		if err != nil {
			fmt.Println(ifURL)
			fmt.Println(err.Error())
			continue
		}
		var iface EthernetInterface
		err = json.Unmarshal(ifData, &iface)
		// Data cannot convert EthernetInterface struct
		if err != nil {
			fmt.Println(ifURL)
			fmt.Println(err.Error())
			continue
		}
		// TODO go routine end
		m.sysEthernetInterface(ch, iface)
	}
	// Systems Ethernet Interfaces end

	// Systems Ethernet Interfaces end

	// storeVolumesLink := store.Volumes.ODataID

	// Set system_health metric

	// biosLink := sys.Bios.ODataID
	// networkIfacesLink := sys.NetworkInterfaces.ODataID
	// biosLink := sys.Bios.ODataID
	// Set Metrics
}

// Get System Data start
// func (m Metrics) Data(url string, c redfish.APIClient) ([]byte, error) {
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
		fmt.Sprintf("%s", s.Id),
		fmt.Sprintf("%d", s.CapacityBytes),
		s.Protocol,
		s.MediaType,
	)
}

func (m Metrics) sysEthernetInterface(ch chan<- prometheus.Metric, iface EthernetInterface) {
	// d, _ := json.MarshalIndent(s, "", "    ")
	// fmt.Println(string(d))
	ch <- prometheus.MustNewConstMetric(
		base.SysEthernetInterface,
		prometheus.GaugeValue,
		iface.PortStatus(),
		iface.MACAddress,
		fmt.Sprintf("%d", iface.SpeedMbps),
	)
}

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
	// Create a copy redfish.Client
	c = *redfish.Client
	fmt.Println(c.Host)
	// var chasCollection
	var chasCollection ChassisCollection
	chasCollectionURL := fmt.Sprintf("%s%s", c.Host, "/redfish/v1/Chassis")
	data, err := redfish.Client.Get(chasCollectionURL)
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
	var chasURL string
	for _, v := range chasCollection.Members {
		if strings.Contains(v.ODataID, "System") {
			chasURL = fmt.Sprintf("%s%s", c.Host, v.ODataID)
			fmt.Println(chasURL)
			break
		}
	}

	// get Chassis Data
	dataChassis, err := c.Get(chasURL)
	// Problem connect to server
	if err != nil {
		return
	}

	var chas Chassis
	// Data cannot convert Chassis struct
	err = json.Unmarshal(dataChassis, &chas)
	if err != nil {
		return
	}
	b, _ := json.MarshalIndent(chas, "", "    ")
	fmt.Println(string(b))
	// m.sysHealth(ch, *redfish.Client)
}

// func (m Metrics) chasPowerControl(ch chan<- prometheus.Metric, c redfish.APIClient) {
// 	data, err := c.Get(c.ChasURL)
// 	// Problem connect to server
// 	if err != nil {
// 		// fmt.Println(err.Error())
// 		// ch <- prometheus.MustNewConstMetric(
// 		// 	base.ChasPower,
// 		// 	prometheus.GaugeValue,
// 		// 	2.0,
// 		// 	"",
// 		// 	"",
// 		// 	"",
// 		// )
// 		return
// 	}
// 	var chas Chassis
// 	err = json.Unmarshal(data, &chas)
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
// 	// ch <- prometheus.MustNewConstMetric(
// 	// 	base.ChasPower,
// 	// 	prometheus.GaugeValue,
// 	// 	sys.StatusToNumber(),
// 	// 	sys.PartNumber,
// 	// 	sys.SKU,
// 	// 	sys.SerialNumber,
// 	// )
// }
