package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

type EthernetInterfaceCollection struct {
	base.Meta
	Description       string `json:"Description"`
	Members           []base.Link
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`
}
type EthernetInterface struct {
	base.Meta
	FullDuplex bool   `json:"FullDuplex"`
	HostName   string `json:"HostName"`
	Id         string `json:"Id"`
	LinkStatus string `json:"LinkStatus"`
	MACAddress string `json:"MACAddress"`
	MTUSize    int    `json:"MTUSize"`
	Name       string `json:"Name"`
	SpeedMbps  int    `json:"SpeedMbps"`
	Status     base.Status

	/*
		{
			"@odata.context": "/redfish/v1/$metadata#EthernetInterface.EthernetInterface",
			"@odata.etag": "W/\"ED874B72\"",
			"@odata.id": "/redfish/v1/Systems/1/EthernetInterfaces/1",
			"@odata.type": "#EthernetInterface.v1_4_1.EthernetInterface",
			"FullDuplex": true,
			"IPv4Addresses": [
				{
					"Address": "42.118.242.149"
				}
			],
			"IPv4StaticAddresses": [],
			"IPv6AddressPolicyTable": [],
			"IPv6Addresses": [
				{
					"Address": "fe80::4adf:37ff:fe8d:1920/64"
				}
			],
			"IPv6StaticAddresses": [],
			"IPv6StaticDefaultGateways": [],
			"Id": "1",
			"InterfaceEnabled": true,
			"LinkStatus": "LinkUp",
			"MACAddress": "48:df:37:8d:19:20",
			"Name": "bond0",
			"NameServers": [],
			"SpeedMbps": 10000,
			"StaticNameServers": [],
			"Status": {
				"Health": "OK",
				"State": "Enabled"
			},
			"UefiDevicePath": "PciRoot(0x8)/Pci(0x0,0x0)/Pci(0x0,0x0)"
		}
	*/
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
