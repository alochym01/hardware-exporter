package dell

import (
	"github.com/alochym01/hardware-exporter/domain/server/base"
)

type ChassisCollection struct {
	base.ChassisCollection
}

type Chassis struct {
	base.Meta
	base.Chassis
	base.Actions
	Assembly         base.Link
	Description      string `json:"Description"`
	Links            ChassisLinks
	Location         ChassisLocation
	Oem              ChassisOEM
	PCIeSlots        base.Link
	PartNumber       string `json:"PartNumber"`
	PhysicalSecurity ChassisPhysicalSecurity
	Sensors          base.Link
	Status           base.Status
	UUID             string `json:"UUID"`
}

func (c Chassis) StatusToNumber() float64 {
	switch c.Status.Health {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	default:
		return 3.0
	}
}
