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

// Describe a description of metrics
func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
	ch <- base.SysStorageState
	ch <- base.SysStorageDisk
	ch <- base.SysEthernetInterface
	ch <- base.ChasPower
}

// Collect return a metric with all desc value and metric value
func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	// System Metrics
	m.SystemsCollector(ch, redfish.ClientDELL)
	// Chassis Metrics
	m.ChassisCollector(ch, redfish.ClientDELL)
}

// NewMetrics return a Metrics struct
func NewMetrics() Metrics {
	return Metrics{}
}

func (m Metrics) SystemsCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// Get server to get metrics
	server := c.Server

	// // get System Data
	sys, err := SetSystemHealthMetric(ch, server)
	if err != nil {
		return
	}
	// Set all Links

	// Systems Storage start
	// Systems Storage Collection Data
	storeCollectionLink := fmt.Sprintf("%s%s", server, sys.Storage.ODataID)
	store, err := SetStorageStatusMetric(ch, server, storeCollectionLink)
	if err != nil {
		return
	}

	// Systems Storage Disk start
	// Set Systems Storage Disk metric
	SetStorageDiskMetric(ch, server, *store)
	// // Systems Storage Disk end

	// Systems Ethernet Interfaces start
	ethIfacesLink := fmt.Sprintf("%s%s", server, sys.EthernetInterfaces.ODataID)
	SetEthernetMetric(ch, server, ethIfacesLink)
	// Systems Ethernet Interfaces end

	// storeVolumesLink := store.Volumes.ODataID

	// Set system_health metric

	// biosLink := sys.Bios.ODataID
	// networkIfacesLink := sys.NetworkInterfaces.ODataID
	// biosLink := sys.Bios.ODataID
	// Set Metrics
}

func SetStorageDiskMetric(ch chan<- prometheus.Metric, server string, store Storage) {
	for _, v := range store.Drives {
		// TODO go routine start
		diskURL := fmt.Sprintf("%s%s", server, v.ODataID)
		// Get Systems Storage Disk Data
		diskData, err := redfish.ClientDELL.Get(diskURL)
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
			// m.sysStorageDisk(ch, disk)
			enclosure := strings.Split(disk.Id, ".")[2]
			id := strings.Split(enclosure, ":")[0]
			ch <- prometheus.MustNewConstMetric(
				base.SysStorageDisk,
				prometheus.GaugeValue,
				disk.PredictedMediaLifeLeftPercent,
				fmt.Sprintf("%s", id),
				fmt.Sprintf("%d", disk.CapacityBytes/1000000000),
				disk.Protocol,
				disk.MediaType,
			)
		}
	}
	return
}

func SetStorageStatusMetric(ch chan<- prometheus.Metric, server string, url string) (*Storage, error) {
	data, err := redfish.ClientDELL.Get(url)
	// Problem connect to server
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return nil, err
	}
	var storeCollection StorageCollection
	err = json.Unmarshal(data, &storeCollection)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return nil, err
	}
	// Systems Storage Collection End

	// Systems Storage Data start
	// Set Systems Storage URL
	var storeURL string
	for _, v := range storeCollection.Members {
		if strings.Contains(v.ODataID, "RAID") {
			// storeURL = fmt.Sprintf("%s%s", redfish.ClientDELL.Host, v.ODataID)
			storeURL = fmt.Sprintf("%s%s", server, v.ODataID)
			break
		}
	}
	// Get Systems Storage Data
	storeData, err := redfish.ClientDELL.Get(storeURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(storeURL)
		fmt.Println(err.Error())
		return nil, err
	}

	var store Storage
	err = json.Unmarshal(storeData, &store)
	// Data cannot convert Storage struct
	if err != nil {
		fmt.Println(storeURL)
		fmt.Println(err.Error())
		return nil, err
	}
	// Set storage_status metric
	ch <- prometheus.MustNewConstMetric(
		base.SysStorageState,
		prometheus.GaugeValue,
		store.StatusToNumber(),
	)
	// m.sysStorageStatus(ch, store) // Systems Storage Data end
	return &store, nil
}

func SetSystemHealthMetric(ch chan<- prometheus.Metric, server string) (*Systems, error) {
	var sysCollection SystemsCollection
	sysCollectionURL := fmt.Sprintf("%s%s", server, "/redfish/v1/Systems")
	data, err := redfish.ClientDELL.Get(sysCollectionURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(sysCollectionURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return nil, err
	}
	err = json.Unmarshal(data, &sysCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(sysCollectionURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return nil, err
	}
	// SystemsCollection end

	// Systems start
	// Set a systems url
	var sysURL string
	for _, v := range sysCollection.Members {
		sysURL = fmt.Sprintf("%s%s", server, v.ODataID)
	}

	// get System Data
	sysData, err := redfish.ClientDELL.Get(sysURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(sysURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return nil, err
	}
	var sys Systems
	err = json.Unmarshal(sysData, &sys)
	// Data cannot convert System struct
	if err != nil {
		fmt.Println(sysURL)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return nil, err
	}
	ch <- prometheus.MustNewConstMetric(
		base.SysState,
		prometheus.GaugeValue,
		sys.StatusToNumber(),
		sys.SKU,
		sys.SerialNumber,
	)
	return &sys, nil
}

func SetEthernetMetric(ch chan<- prometheus.Metric, server string, url string) {
	// Systems Ethernet Interfaces Collection
	data, err := redfish.ClientDELL.Get(url)
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	var ethIf EthernetInterfaceCollection
	err = json.Unmarshal(data, &ethIf)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	// Systems Ethernet Interfaces Collection end
	// Systems Ethernet Interfaces start
	for _, v := range ethIf.Members {
		// TODO go routine start
		ifURL := fmt.Sprintf("%s%s", server, v.ODataID)
		// Get Ethernet Interfaces Data
		ifData, err := redfish.ClientDELL.Get(ifURL)
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
		ch <- prometheus.MustNewConstMetric(
			base.SysEthernetInterface,
			prometheus.GaugeValue,
			iface.PortStatus(),
			iface.Id,
			iface.MACAddress,
			fmt.Sprintf("%d", iface.SpeedMbps),
		)
	}
}

func (m Metrics) ChassisCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// Get server to get metric
	server := c.Server

	// Chassis Collection start
	chasCollectionURL := fmt.Sprintf("%s%s", server, "/redfish/v1/Chassis")
	chas, err := GetChassis(chasCollectionURL, server)
	if err != nil {
		return
	}
	// Chassis Collection end

	// Set PowerControl Link
	chasPowerLink := fmt.Sprintf("%s%s", server, chas.Power.ODataID)
	SetPowerMetrics(ch, chasPowerLink)

	// Set Thermal Link
	// chasThermalLink := fmt.Sprintf("%s%s", Host, chas.Thermal.ODataID)
}

func GetChassis(url string, server string) (*Chassis, error) {
	var chasCollection ChassisCollection
	dataCollection, err := redfish.ClientDELL.Get(url)
	// Problem connect to server
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(dataCollection, &chasCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return nil, err
	}
	// b, _ := json.MarshalIndent(chas, "", "    ")
	// fmt.Println(string(b))

	// Chassis start
	// Set a chassis url
	var chasURL string
	for _, v := range chasCollection.Members {
		if strings.Contains(v.ODataID, "System") {
			chasURL = fmt.Sprintf("%s%s", server, v.ODataID)
			break
		}
	}

	// get Chassis Data
	dataChassis, err := redfish.ClientDELL.Get(chasURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(chasURL)
		fmt.Println(err.Error())
		return nil, err
	}

	var chas Chassis
	// Data cannot convert Chassis struct
	err = json.Unmarshal(dataChassis, &chas)
	if err != nil {
		fmt.Println(chasURL)
		fmt.Println(err.Error())
		return nil, err
	}

	// b, _ := json.MarshalIndent(chas, "", "    ")
	// fmt.Println(string(b))
	return &chas, nil
}

func SetPowerMetrics(ch chan<- prometheus.Metric, url string) {
	data, err := redfish.ClientDELL.Get(url)
	// Problem connect to server
	if err != nil {
		return
	}
	var power PowerControl
	err = json.Unmarshal(data, &power)
	// Data cannot convert PowerControl struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	// b, _ := json.MarshalIndent(power, "", "    ")
	// fmt.Println(string(b))
	// Everything is ok
	for _, v := range power.PowerControl {
		ch <- prometheus.MustNewConstMetric(
			base.ChasPower,
			prometheus.GaugeValue,
			v.PowerConsumedWatts,
		)
	}
}
