package hpe

import (
	"encoding/json"
	"fmt"

	"github.com/alochym01/hardware-exporter/domain/server/base"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
}

func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
	ch <- base.SysStorageState
	ch <- base.SysStorageDisk
	ch <- base.SysEthernetInterface
	ch <- base.ChasPower
}
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
	// Get server to get metrics
	server := c.Server
	// // get System Data
	sys, err := SetSystemHealthMetric(ch, server)
	if err != nil {
		return
	}
	// Systems Storage start
	// Systems Storage Collection Data
	storeCollectionLink := fmt.Sprintf("%s%s", server, sys.Oem.HPE.Links.SmartStorage.ODataID)
	fmt.Println(storeCollectionLink)
	// store, err := SetStorageStatusMetric(ch, server, storeCollectionLink)
	store, err := SetStorageStatusMetric(ch, server, storeCollectionLink)
	if err != nil {
		return
	}

	// Systems Storage Disk start
	// Set Systems Storage Disk metric
	SetStorageDiskMetric(ch, server, store)
	// Systems Storage Disk end
	// Systems Ethernet Interfaces start
	ethIfacesLink := fmt.Sprintf("%s%s", server, sys.EthernetInterfaces.ODataID)
	SetEthernetMetric(ch, server, ethIfacesLink)
	// Systems Ethernet Interfaces end

	// SetEthernetMetric(ch chan<- prometheus.Metric, server string, url string)
}
func SetStorageDiskMetric(ch chan<- prometheus.Metric, server string, store *StorageArrayController) {
	var diskCollection SmartStorageDiskDriveCollection
	diskCollectionURL := fmt.Sprintf("%s%s", server, store.Links.PhysicalDrives.ODataID)
	diskCollectionData, err := redfish.Client.Get(diskCollectionURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(diskCollectionURL)
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(diskCollectionData, &diskCollection)
	// Data cannot convert SmartStorageDiskDriveCollection struct
	if err != nil {
		fmt.Println(diskCollectionURL)
		fmt.Println(err.Error())
		return
	}

	for _, v := range diskCollection.Members {
		// TODO go routine start
		diskURL := fmt.Sprintf("%s%s", server, v.ODataID)
		// Get Systems Storage Disk Data
		diskData, err := redfish.Client.Get(diskURL)
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
		// d, _ := json.MarshalIndent(disk, "", "    ")
		// fmt.Println(string(d))
		// TODO go routine end
		// Check Disk is SSD
		if disk.SSDEnduranceUtilizationPercentage > 0 {
			// m.sysStorageDisk(ch, disk)
			ch <- prometheus.MustNewConstMetric(
				base.SysStorageDisk,
				prometheus.GaugeValue,
				(100.0 - disk.SSDEnduranceUtilizationPercentage),
				fmt.Sprintf("%s", disk.Id),
				fmt.Sprintf("%d", disk.CapacityMiB),
				disk.InterfaceType,
				disk.MediaType,
			)
		}
	}
	return
}
func SetSystemHealthMetric(ch chan<- prometheus.Metric, server string) (*Systems, error) {
	var sysCollection SystemsCollection
	sysCollectionURL := fmt.Sprintf("%s%s", server, "/redfish/v1/Systems")
	data, err := redfish.Client.Get(sysCollectionURL)
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
	sysData, err := redfish.Client.Get(sysURL)
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
func SetStorageStatusMetric(ch chan<- prometheus.Metric, server string, url string) (*StorageArrayController, error) {
	data, err := redfish.Client.Get(url)
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
	// Set storage_status metric
	ch <- prometheus.MustNewConstMetric(
		base.SysStorageState,
		prometheus.GaugeValue,
		storeCollection.StatusToNumber(),
	)

	// Systems Storage Data start
	// Set Systems Storage Collection URL
	var arrayControllerCollectionURL string
	arrayControllerCollectionURL = fmt.Sprintf("%s%s", server, storeCollection.Links.ArrayControllers.ODataID)
	// Get Systems Storage Data
	arrayControllerCollectionData, err := redfish.Client.Get(arrayControllerCollectionURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(arrayControllerCollectionURL)
		fmt.Println(err.Error())
		return nil, err
	}
	var storeArrControllerCollection StorageArrayControllerCollection
	err = json.Unmarshal(arrayControllerCollectionData, &storeArrControllerCollection)
	// d, _ := json.MarshalIndent(storeArrControllerCollection, "", "    ")
	// fmt.Println(string(d))
	// Data cannot convert Storage struct
	if err != nil {
		fmt.Println(arrayControllerCollectionURL)
		fmt.Println(err.Error())
		return nil, err
	}
	// Systems Storage Collection End

	// Systems Storage ArrayController Start
	// Set Systems Storage ArrayController URL
	var arrayControllerURL string
	arrayControllerURL = fmt.Sprintf("%s%s", server, storeArrControllerCollection.Members[0].ODataID)
	// Get Systems Storage Data
	arrayControllerData, err := redfish.Client.Get(arrayControllerURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(arrayControllerURL)
		fmt.Println(err.Error())
		return nil, err
	}

	var storeArrController StorageArrayController
	err = json.Unmarshal(arrayControllerData, &storeArrController)
	// d, _ := json.MarshalIndent(storeArrController, "", "    ")
	// fmt.Println(string(d))
	// Data cannot convert Storage struct
	if err != nil {
		fmt.Println(arrayControllerURL)
		fmt.Println(err.Error())
		return nil, err
	}

	// m.sysStorageStatus(ch, store) // Systems Storage Data end
	return &storeArrController, nil
}

func SetEthernetMetric(ch chan<- prometheus.Metric, server string, url string) {
	// Systems Ethernet Interfaces Collection
	data, err := redfish.Client.Get(url)
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
		ifData, err := redfish.Client.Get(ifURL)
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
	dataCollection, err := redfish.Client.Get(url)
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
		chasURL = fmt.Sprintf("%s%s", server, v.ODataID)
	}

	// get Chassis Data
	dataChassis, err := redfish.Client.Get(chasURL)
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
	data, err := redfish.Client.Get(url)
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
