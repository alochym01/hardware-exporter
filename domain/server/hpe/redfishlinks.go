package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

type RedfishLinksInstances struct {
	ODataID     string   `json:"@odata.id"`
	OdataType   string   `json:"@odata.type"`
	ETag        string   `json:"ETag"`
	HttpMethods []string `json:"HttpMethods"`
	/*{
	    "@odata.id": "/redfish/v1/Systems",
	    "@odata.type": "#ComputerSystemCollection.ComputerSystemCollection",
	    "ETag": "W/\"AA6D42B0\"",
	    "HttpMethods": [
	        "GET",
	        "HEAD"
	    ]
	}*/
}
type RedfishLinks struct {
	base.Meta
	Id        string `json:"Id"`
	OdataType string `json:"@odata.type"`
	Instances []RedfishLinksInstances
}
