package dell

import (
	"github.com/alochym01/hardware-exporter/domain/server/base"
)

type EthernetInterfaceCollection struct {
	base.Meta
	Description       string `json:"Description"`
	Members           []base.Link
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`
}

type EthernetInterface struct {
	base.Meta
	AutoNeg     bool   `json:"AutoNeg"`
	Description string `json:"Description"`
	FQDN        string `json:"FQDN"`
	FullDuplex  bool   `json:"FullDuplex"`
	HostName    string `json:"HostName"`
	Id          string `json:"Id"`
	LinkStatus  string `json:"LinkStatus"`
	MACAddress  string `json:"MACAddress"`
	MTUSize     int    `json:"MTUSize"`
	Name        string `json:"Name"`
	SpeedMbps   int    `json:"SpeedMbps"`
	Status      base.Status
	// Links       SystemsEthernetInterfaceLinks
	// InterfaceEnabled bool   `json:"InterfaceEnabled"`
	// UefiDevicePath   string `json:"UefiDevicePath"`
	// IPv4Addresses       []string `json:"IPv4Addresses"`
	// IPv6Addresses       []string `json:"IPv6Addresses"`
	// PermanentMACAddress string `json:"PermanentMACAddress"`
}

func (e EthernetInterface) PortStatus() float64 {
	switch e.SpeedMbps {
	case 0:
		return 2.0
	default:
		return 0.0
	}
	// switch e.LinkStatus {
	// case "LinkUp":
	// 	return 0.0
	// case "LinkDown":
	// 	return 2.0
	// default:
	// 	return 3.0
	// }
}
