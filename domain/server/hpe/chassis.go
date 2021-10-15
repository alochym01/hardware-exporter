package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

type ChassisCollection struct {
	base.ChassisCollection
	ODataEtag string `json:"@odata.etag"`
}

type Chassis struct {
	base.Chassis
	Links     base.ChassisLinks
	ODataEtag string `json:"@odata.etag"` // HPE
	Oem       ChassisOEM
}
