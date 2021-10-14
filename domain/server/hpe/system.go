package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

type Systems struct {
	base.Meta
	ODataEtag string `json:"@odata.etag"` // HPE
	base.Actions
	AssetTag           string `json:"AssetTag"`
	Boot               SystemsBoot
	Bios               base.Link
	BiosVersion        string `json:"BiosVersion"`
	EthernetInterfaces base.Link
	HostName           string `json:"HostName"`
	Id                 string `json:"Id"`
	IndicatorLED       string `json:"IndicatorLED"`
	Links              SystemsLinks
	LogServices        base.Link
	Manufacturer       string `json:"Manufacturer"`
	Memory             base.Link
	MemoryDomains      base.Link
	MemorySummary      SystemsMemorySummary
	Model              string `json:"Model"`
	Name               string `json:"Name"`
	NetworkInterfaces  base.Link
	Oem                SystemsOem
	PowerState         string `json:"PowerState"`
	ProcessorSummary   SystemsProcessorSummary
	Processors         base.Link
	SKU                string `json:"SKU"`
	SecureBoot         base.Link
	SerialNumber       string `json:"SerialNumber"`
	Status             base.Status
	Storage            base.Link
	SystemType         string `json:"SystemType"`
	TrustedModules     []SystemsTrustedModules
	UUID               string `json:"UUID"`
}
