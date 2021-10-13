package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

// ChassisActions start
type ChassisActions struct {
	ChassisReset ChassisActionsChassisReset `json:"Chassis.Reset"`
}

type ChassisActionsChassisReset struct {
	Target string `json:"target"`
	// ResetType []string `json:"ResetTypeRedfish.AllowableValues"`
}

// ChassisActions end

// ChassisLinks start
type ChassisLinks struct {
	base.ChassisLinks
	ComputerSystemsOdataCount   int         `json:"ComputerSystems@odata.count"`
	Contains                    []base.Link `json:"Contains"`
	ContainsOdataCount          int         `json:"Contains@odata.count"`
	CooledBy                    []base.Link
	CooledByOdataCount          int      `json:"CooledBy@odata.count"`
	Drives                      []string `json:"Drives"`
	DrivesOdataCount            int      `json:"Drives@odata.count"`
	ManagedByOdataCount         int      `json:"ManagedBy@odata.count"`
	ManagersInChassis           []base.Link
	ManagersInChassisOdataCount int `json:"ManagersInChassis@odata.count"`
	PCIeDevices                 []base.Link
	PCIeDevicesOdataCount       int `json:"PCIeDevices@odata.count"`
	Processors                  []base.Link
	ProcessorsOdataCount        int `json:"Processors@odata.count"`
	PoweredBy                   []base.Link
	PoweredByOdataCount         int `json:"PoweredBy@odata.count"`
	Storage                     []base.Link
	StorageOdataCount           int `json:"Storage@odata.count"`
}

// ChassisLocation start
type ChassisLocation struct {
	Info          string `json:"Info"`
	InfoFormat    string `json:"InfoFormat"`
	Placement     ChassisLocationPlacement
	PostalAddress ChassisLocationPostalAddress
}
type ChassisLocationPostalAddress struct {
	Building string `json:"Building"`
	Room     string `json:"Room"`
}
type ChassisLocationPlacement struct {
	Rack string `json:"Rack"`
	Row  string `json:"Row"`
}

// ChassisLocation end

// ChassisOEM start
type ChassisOEM struct {
	Dell ChassisOEMDell
}
type ChassisOEMDell struct {
	DellChassis ChassisOEMDellDellChassis
}

type ChassisOEMDellDellChassis struct {
	base.Meta
	CanBeFRUed bool `json:"CanBeFRUed"`
	Links      ChassisOEMDellDellChassisLinks
	SystemID   int `json:"SystemID"`
}

type ChassisOEMDellDellChassisLinks struct {
	ComputerSystem base.Link
}

// ChassisOEM end

type ChassisPhysicalSecurity struct {
	IntrusionSensor       string `json:"IntrusionSensor"`
	IntrusionSensorNumber int    `json:"IntrusionSensorNumber"`
	IntrusionSensorReArm  string `json:"IntrusionSensorReArm"`
}
type ChassisStatus struct {
	base.Status
	HealthRollup string `json:"HealthRollup"`
}
