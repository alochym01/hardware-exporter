package hpe

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alochym01/hardware-exporter/domain/server/base"
	"github.com/alochym01/hardware-exporter/storage/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

var rfLinks *RedfishLinks

type MetricsV1 struct {
}

// NewMetrics return a Metrics struct
func NewMetricsV1() MetricsV1 {
	return MetricsV1{}
}
func (m MetricsV1) Describe(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
	ch <- base.SysStorageState
	ch <- base.SysStorageDisk
	ch <- base.SysEthernetInterface
	ch <- base.ChasPower
}

func (m MetricsV1) Collect(ch chan<- prometheus.Metric) {
	// clientHPE:=
	server := redfish.ClientHPE.Server
	// Get all Redfish Link
	redfishLinks := fmt.Sprintf("%s%s", server, "/redfish/v1/resourcedirectory")
	data, err := redfish.ClientHPE.Get(redfishLinks)
	if err != nil {
		fmt.Println(redfishLinks)
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}

	err = json.Unmarshal(data, &rfLinks)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(redfishLinks)
		fmt.Println(err.Error())
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return
	}

	// System Metrics
	m.SystemsCollector(ch, redfish.ClientHPE, server)
	// Chassis Metrics
	m.ChassisCollector(ch, redfish.ClientHPE, server)
}

func (m MetricsV1) SystemsCollector(ch chan<- prometheus.Metric, c redfish.APIClient, server string) {
	SetSystemHealthMetricsV1(ch, rfLinks, server)
	// DiskDrive.
	SetStorageDiskMetricsV1(ch, rfLinks, server)
	// EthernetInterface.
	SetEthernetMetricsV1(ch, rfLinks, server)
}

// func SetSystemHealthMetricsV1(ch chan<- prometheus.Metric, server string) (*Systems, error) {
func SetSystemHealthMetricsV1(ch chan<- prometheus.Metric, sysLink *RedfishLinks, server string) {
	// get System Data
	sysURL := findObject(sysLink.Instances, "ComputerSystem.", server)

	for _, v := range sysURL {
		if v == "" {
			fmt.Println(v)
			fmt.Println("Not Found")
			ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return
		}
		sysData, err := redfish.ClientHPE.Get(v)
		// Problem connect to server
		if err != nil {
			fmt.Println(v)
			fmt.Println(err.Error())
			ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return
		}
		var sys Systems
		err = json.Unmarshal(sysData, &sys)
		// Data cannot convert System struct
		if err != nil {
			fmt.Println(v)
			fmt.Println(err.Error())
			ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return
		}
		// fmt.Println("ALOCHYM", v)
		ch <- prometheus.MustNewConstMetric(
			base.SysState,
			prometheus.GaugeValue,
			sys.StatusToNumber(),
			sys.SKU,
			sys.SerialNumber,
		)
	}
	return
}

// func SetStorageDiskMetricsV1(ch chan<- prometheus.Metric, server string, store *StorageArrayController) {
func SetStorageDiskMetricsV1(ch chan<- prometheus.Metric, sysLink *RedfishLinks, server string) {
	diskURLs := findObject(sysLink.Instances, "DiskDrive.", server)
	// Using goroutine
	diskalochym := make(chan []byte)

	// Using go routine
	for _, diskURL := range diskURLs {
		go redfish.ClientHPE.GetUseGoRoutine(diskURL, diskalochym)
	}

	// TODO go routine end
	for range diskURLs {
		var disk StorageDisk
		diskData := <-diskalochym
		json.Unmarshal(diskData, &disk)
		// fmt.Println(disk.ODataID)
		// Check Disk is SSD
		if disk.SSDEnduranceUtilizationPercentage > 0 {
			ch <- prometheus.MustNewConstMetric(
				base.SysStorageDisk,
				prometheus.GaugeValue,
				(100.0 - disk.SSDEnduranceUtilizationPercentage),
				fmt.Sprintf("%s", disk.Id),
				fmt.Sprintf("%d", disk.CapacityMiB/1000),
				disk.InterfaceType,
				disk.MediaType,
			)
		}
	}

	return
}

func SetEthernetMetricsV1(ch chan<- prometheus.Metric, sysLink *RedfishLinks, server string) {
	// Systems Ethernet Interfaces Collection
	ifURLs := findObject(sysLink.Instances, "EthernetInterface.", server)
	// // Using goroutine
	// // TODO go routine start
	// Systems Ethernet Interfaces Collection end
	// Systems Ethernet Interfaces start
	ifalochym := make(chan []byte)

	// Using go routine
	for _, ifURL := range ifURLs {
		go redfish.ClientHPE.GetUseGoRoutine(ifURL, ifalochym)
	}

	// Get Ethernet Interfaces Data
	for range ifURLs {
		var iface EthernetInterface
		ifData := <-ifalochym
		json.Unmarshal(ifData, &iface)
		ch <- prometheus.MustNewConstMetric(
			base.SysEthernetInterface,
			prometheus.GaugeValue,
			iface.PortStatus(),
			iface.MACAddress,
			fmt.Sprintf("%d", iface.SpeedMbps),
		)
	}
	// TODO go routine end
	return
}
func (m MetricsV1) ChassisCollector(ch chan<- prometheus.Metric, c redfish.APIClient, server string) {
	// Set PowerControl Link
	// chasPowerLink := fmt.Sprintf("%s%s", server, chas.Power.ODataID)
	SetPowerMetricsV1(ch, rfLinks, server)

	// Set Thermal Link
	// chasThermalLink := fmt.Sprintf("%s%s", Host, chas.Thermal.ODataID)
}

func SetPowerMetricsV1(ch chan<- prometheus.Metric, sysLink *RedfishLinks, server string) {
	// Systems Ethernet Interfaces Collection
	ifPowers := findObject(sysLink.Instances, "Power.", server)
	ifpower := make(chan []byte)

	// Using go routine
	for _, ifPower := range ifPowers {
		go redfish.ClientHPE.GetUseGoRoutine(ifPower, ifpower)
	}

	for range ifPowers {
		var power PowerControl
		data := <-ifpower
		err := json.Unmarshal(data, &power)
		// Data cannot convert PowerControl struct
		if err != nil {
			// fmt.Println(ifPowers[0])
			fmt.Println(err.Error())
			return
		}
		for _, v := range power.PowerControl {
			ch <- prometheus.MustNewConstMetric(
				base.ChasPower,
				prometheus.GaugeValue,
				v.PowerConsumedWatts,
			)
		}
	}
}

func findObject(ob []RedfishLinksInstances, obType string, server string) []string {
	var l []string
	for i := range ob {
		if strings.Contains(ob[i].OdataType, obType) && obType == "ComputerSystem." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			l = append(l, url)
			return l
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "Power." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			l = append(l, url)
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "DiskDrive." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			l = append(l, url)
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "EthernetInterface." {
			if strings.Contains(ob[i].ODataID, "Systems") {
				url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
				// fmt.Println(url)
				l = append(l, url)
			}
		}
	}
	return l
}
