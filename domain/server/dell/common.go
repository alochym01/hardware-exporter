package dell

type Link struct {
	ODataID string `json:"@odata.id"`
}

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
	ComputerSystems             []Link `json:"ComputerSystems"`
	ComputerSystemsOdataCount   int    `json:"ComputerSystems@odata.count"`
	Contains                    []Link `json:"Contains"`
	ContainsOdataCount          int    `json:"Contains@odata.count"`
	CooledBy                    []Link
	CooledByOdataCount          int      `json:"CooledBy@odata.count"`
	Drives                      []string `json:"Drives"`
	DrivesOdataCount            int      `json:"Drives@odata.count"`
	ManagedBy                   []Link
	ManagedByOdataCount         int `json:"ManagedBy@odata.count"`
	ManagersInChassis           []Link
	ManagersInChassisOdataCount int `json:"ManagersInChassis@odata.count"`
	PCIeDevices                 []Link
	PCIeDevicesOdataCount       int `json:"PCIeDevices@odata.count"`
	PoweredBy                   []Link
	PoweredByOdataCount         int `json:"PoweredBy@odata.count"`
	Storage                     []Link
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
	ODataContext string `json:"@odata.context"`
	ODataID      string `json:"@odata.id"`
	ODataType    string `json:"@odata.type"`
	CanBeFRUed   bool   `json:"CanBeFRUed"`
	Links        ChassisOEMDellDellChassisLinks
	SystemID     int `json:"SystemID"`
}

type ChassisOEMDellDellChassisLinks struct {
	ComputerSystem Link
}

// ChassisOEM end

type ChassisPhysicalSecurity struct {
	IntrusionSensor       string `json:"IntrusionSensor"`
	IntrusionSensorNumber int    `json:"IntrusionSensorNumber"`
	IntrusionSensorReArm  string `json:"IntrusionSensorReArm"`
}
type ChassisStatus struct {
	Health       string `json:"Health"`
	HealthRollup string `json:"HealthRollup"`
	State        string `json:"State"`
}
