package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

type Systems struct {
	base.Meta
	base.Actions
	AssetTag               string `json:"AssetTag"`
	Bios                   base.Link
	BiosVersion            string `json:"BiosVersion"`
	Boot                   SystemsBoot
	Description            string `json:"Description"`
	EthernetInterfaces     base.Link
	HostName               string `json:"HostName"`
	HostWatchdogTimer      SystemHostWatchdogTimer
	HostingRoles           []string `json:"HostingRoles"`
	HostingRolesOdataCount int      `json:"HostingRoles@odata.count"`
	Id                     string   `json:"Id"`
	IndicatorLED           string   `json:"IndicatorLED"`
	Links                  SystemsLinks
	Manufacturer           string `json:"Manufacturer"`
	Memory                 base.Link
	MemorySummary          SystemsMemorySummary
	Model                  string `json:"Model"`
	Name                   string `json:"Name"`
	NetworkInterfaces      base.Link
	Oem                    SystemsOEM
	PCIeDevices            []base.Link
	PCIeDevicesOdataCount  int    `json:"PCIeDevices@odata.count"`
	PartNumber             string `json:"PartNumber"`
	PowerState             string `json:"PowerState"`
	ProcessorSummary       SystemsProcessorSummary
	Processors             base.Link
	SKU                    string `json:"SKU"`
	SecureBoot             base.Link
	SerialNumber           string `json:"SerialNumber"`
	SimpleStorage          base.Link
	Status                 base.Status
	Storage                base.Link
	SystemType             string `json:"SystemType"`
	TrustedModules         []SystemTrustedModules
	UUID                   string `json:"UUID"`
}
