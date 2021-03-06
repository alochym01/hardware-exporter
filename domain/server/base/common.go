package base

type Link struct {
	ODataID string `json:"@odata.id"`
}

type ComputerSystemReset struct {
	ResetType []string `json:"ResetType@Redfish.AllowableValues"`
	Target    string   `json:"Target"`
}
type Action struct {
	ComputerSystem ComputerSystemReset `json:"#ComputerSystem.Reset"`
}
type Actions struct {
	Actions Action `json:"Actions"`
}
type Meta struct {
	ODataContext string `json:"@odata.context"`
	ODataID      string `json:"@odata.id"`
	ODataType    string `json:"@odata.type"`
}
type Chassis struct {
	Meta
	AssetTag        string `json:"AssetTag"`
	ChassisType     string `json:"ChassisType"`
	Id              string `json:"Id"`
	IndicatorLED    string `json:"IndicatorLED"`
	Manufacturer    string `json:"Manufacturer"`
	Model           string `json:"Model"`
	Name            string `json:"Name"`
	NetworkAdapters Link
	Power           Link
	PowerState      string `json:"PowerState"`
	SKU             string `json:"SKU"`
	SerialNumber    string `json:"SerialNumber"`
	Status          Status
	Thermal         Link
}

// Working for Dell and Hpe Server
type ChassisCollection struct {
	Meta
	Description       string `json:"Description"`
	Members           []Link
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`
}

// ChassisLinks start
type ChassisLinks struct {
	ComputerSystems []Link `json:"ComputerSystems"`
	ManagedBy       []Link
}

// ChassisLinks end

type HealthStatus struct {
	Health string `json:"Health"`
}
type Status struct {
	HealthStatus
	StateStatus
	HealthRollupStatus
}
type StateStatus struct {
	State string `json:"State"`
}

type HealthRollupStatus struct {
	HealthRollup string `json:"HealthRollup"`
}
