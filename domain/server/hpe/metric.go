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
	d, _ := json.MarshalIndent(diskCollectionData, "", "    ")
	fmt.Println(string(d))

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
		// TODO go routine end
		// Check Disk is SSD
		if disk.SSDEnduranceUtilizationPercentage > 0 {
			// m.sysStorageDisk(ch, disk)
			ch <- prometheus.MustNewConstMetric(
				base.SysStorageDisk,
				prometheus.GaugeValue,
				disk.SSDEnduranceUtilizationPercentage,
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
	// arrayControllerURL = fmt.Sprintf("%s%s", server, storeCollection.Links.ArrayControllers.ODataID)
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

func (m Metrics) ChassisCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// Get server to get metrics
	// server := c.Server
}
