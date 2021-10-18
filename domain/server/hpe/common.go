package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

// ChassisOEM start
type ChassisOEM struct {
	Hpe ChassisOEMHpe
}
type ChassisOEMHpe struct {
	ODataContext string `json:"@odata.context"`
	// ODataID      string `json:"@odata.id"` // DELL
	ODataType                 string `json:"@odata.type"`
	Actions                   ChassisOEMHpeActions
	Firmware                  ChassisOEMHpeFirmware
	Links                     ChassisOEMHpeLinks
	MCTPEnabledOnServer       bool `json:"MCTPEnabledOnServer"`
	SmartStorageBattery       []ChassisOEMHpeSmartStorageBattery
	SystemMaintenanceSwitches ChassisOEMHpeSystemMaintenanceSwitches
}

// ChassisOEMHpeActions start
type ChassisOEMHpeActions struct {
	HpeServerChassisDisableMCTPOnServer ActionsTarget
	HpeServerChassisFactoryResetMCTP    ActionsTarget
}

type ActionsTarget struct {
	Target string `json:"target"`
}

// ChassisOEMHpeActions end

// ChassisOEMHpeFirmware start
type ChassisOEMHpeFirmware struct {
	PlatformDefinitionTable             FirmwareItem
	PowerManagementController           FirmwareItem
	PowerManagementControllerBootloader PowerManagementControllerBootloader
	SPSFirmwareVersionData              FirmwareItem
	SystemProgrammableLogicDevice       FirmwareItem
}

type PowerManagementControllerBootloader struct {
	Current FirmwareCurrentFamily
}

type FirmwareCurrentFamily struct {
	Family        string `json:"Family"`
	VersionString string `json:"VersionString"`
}
type FirmwareItem struct {
	Current FirmwareCurrent
}

type FirmwareCurrent struct {
	VersionString string `json:"VersionString"`
}

// ChassisOEMHpeFirmware end

// ChassisOEMHpeLinks start
type ChassisOEMHpeLinks struct {
	Devices ChassisOEMHpeLinksDevices
}

type ChassisOEMHpeLinksDevices struct {
	base.Link
}

// ChassisOEMHpeLinks end

type ChassisOEMHpeSmartStorageBattery struct {
	ChargeLevelPercent         int    `json:"ChargeLevelPercent"`
	FirmwareVersion            string `json:"FirmwareVersion"`
	Index                      int    `json:"Index"`
	MaximumCapWatts            int    `json:"MaximumCapWatts"`
	Model                      string `json:"Model"`
	ProductName                string `json:"ProductName"`
	RemainingChargeTimeSeconds int    `json:"RemainingChargeTimeSeconds"`
	SerialNumber               string `json:"SerialNumber"`
	SparePartNumber            string `json:"SparePartNumber"`
	Status                     base.Status
}
type ChassisOEMHpeSystemMaintenanceSwitches struct {
	Sw1  string `json:"Sw1"`
	Sw2  string `json:"Sw2"`
	Sw3  string `json:"Sw3"`
	Sw4  string `json:"Sw4"`
	Sw5  string `json:"Sw5"`
	Sw6  string `json:"Sw6"`
	Sw7  string `json:"Sw7"`
	Sw8  string `json:"Sw8"`
	Sw9  string `json:"Sw9"`
	Sw10 string `json:"Sw10"`
	Sw11 string `json:"Sw11"`
	Sw12 string `json:"Sw12"`
}

// ChassisOEM end

// Systems
// SystemsTrustedModules start
type SystemsTrustedModulesOEMHPE struct {
	ODataContext string `json:"@odata.context"`
	ODataType    string `json:"@odata.type"`
}
type SystemsTrustedModulesOEM struct {
	HPE SystemsTrustedModulesOEMHPE
}

type SystemsTrustedModules struct {
	OEM    SystemsTrustedModulesOEM
	Status base.StateStatus
}

// SystemsTrustedModules end

type SystemsOemHPEPowerSupplies struct {
	PowerSuppliesMismatch bool `json:"PowerSuppliesMismatch"`
	SystemsOemHPEStatus
}
type SystemsProcessorSummary struct {
	Count  int    `json:"Count"`
	Model  string `json:"Model"`
	Status base.HealthRollupStatus
}

type SystemsOemHPEStatus struct {
	Status base.HealthStatus
}
type SystemsOemHPEAggregateHealthStatus struct {
	AgentlessManagementService string `json:"AgentlessManagementService"`
	BiosOrHardwareHealth       SystemsOemHPEStatus
	FanRedundancy              string `json:"FanRedundancy"`
	Fans                       SystemsOemHPEStatus
	Memory                     SystemsOemHPEStatus
	Network                    SystemsOemHPEStatus
	PowerSupplies              SystemsOemHPEPowerSupplies
	PowerSupplyRedundancy      string `json:"PowerSupplyRedundancy"`
	Processors                 SystemsOemHPEStatus
	SmartStorageBattery        SystemsOemHPEStatus
	Storage                    SystemsOemHPEStatus
	Temperatures               SystemsOemHPEStatus
}
type SystemsOemHPEUserDataEraseComponentStatus struct {
}
type SystemsOemHPESystemUsage struct {
}
type SystemsOemHPESystemROMAndiLOEraseComponentStatus struct {
}
type SystemsOemHPESMBIOS struct {
}
type SystemsOemHPEProcessorJitterControl struct {
}
type SystemsOemHPEPowerRegulatorModesSupported struct {
}
type SystemsOemHPEHostOS struct {
}
type SystemsOemHPEDeviceDiscoveryComplete struct {
}
type SystemsOemHPEBios struct {
}
type SystemsOemHPELinks struct {
	EthernetInterfaces         base.Link
	NetworkAdapters            base.Link
	PCIDevices                 base.Link
	PCISlots                   base.Link
	SmartStorage               base.Link
	USBDevices                 base.Link
	USBPorts                   base.Link
	WorkloadPerformanceAdvisor base.Link
}
type SystemsOemHPE struct {
	AggregateHealthStatus           SystemsOemHPEAggregateHealthStatus
	Bios                            SystemsOemHPEBios
	CurrentPowerOnTimeSeconds       int `json:"CurrentPowerOnTimeSeconds"`
	DeviceDiscoveryComplete         SystemsOemHPEDeviceDiscoveryComplete
	ElapsedEraseTimeInMinutes       int    `json:"ElapsedEraseTimeInMinutes"`
	EndOfPostDelaySeconds           string `json:"EndOfPostDelaySeconds"`
	EstimatedEraseTimeInMinutes     int    `json:"EstimatedEraseTimeInMinutes"`
	HostOS                          SystemsOemHPEHostOS
	IntelligentProvisioningAlwaysOn bool   `json:"IntelligentProvisioningAlwaysOn"`
	IntelligentProvisioningIndex    int    `json:"IntelligentProvisioningIndex"`
	IntelligentProvisioningLocation string `json:"IntelligentProvisioningLocation"`
	IntelligentProvisioningVersion  string `json:"IntelligentProvisioningVersion"`
	IsColdBooting                   bool   `json:"IsColdBooting"`
	Links                           SystemsOemHPELinks
	PCAPartNumber                   string   `json:"PCAPartNumber"`
	PCASerialNumber                 string   `json:"PCASerialNumber"`
	PostDiscoveryCompleteTimeStamp  string   `json:"PostDiscoveryCompleteTimeStamp"`
	PostDiscoveryMode               string   `json:"PostDiscoveryMode"`
	PostMode                        string   `json:"PostMode"`
	PostState                       string   `json:"PostState"`
	PowerAllocationLimit            int      `json:"PowerAllocationLimit"`
	PowerAutoOn                     string   `json:"PowerAutoOn"`
	PowerOnDelay                    string   `json:"PowerOnDelay"`
	PowerOnMinutes                  int      `json:"PowerOnMinutes"`
	PowerRegulatorMode              string   `json:"PowerRegulatorMode"`
	PowerRegulatorModesSupported    []string `json:"PowerRegulatorModesSupported"`
	// PowerRegulatorModesSupported        []SystemsOemHPEPowerRegulatorModesSupported
	ProcessorJitterControl              SystemsOemHPEProcessorJitterControl
	SMBIOS                              SystemsOemHPESMBIOS
	ServerFQDN                          string `json:"ServerFQDN"`
	SmartStorageConfig                  []base.Link
	SystemROMAndiLOEraseComponentStatus SystemsOemHPESystemROMAndiLOEraseComponentStatus
	SystemROMAndiLOEraseStatus          string `json:"SystemROMAndiLOEraseStatus"`
	SystemUsage                         SystemsOemHPESystemUsage
	UserDataEraseComponentStatus        SystemsOemHPEUserDataEraseComponentStatus
	UserDataEraseStatus                 string `json:"UserDataEraseStatus"`
	VirtualProfile                      string `json:"VirtualProfile"`
}
type SystemsOem struct {
	// TODO take time to parse
	HPE SystemsOemHPE
}

type SystemsMemorySummary struct {
	Status                         base.HealthRollupStatus
	TotalSystemMemoryGiB           int `json:"TotalSystemMemoryGiB"`
	TotalSystemPersistentMemoryGiB int `json:"TotalSystemPersistentMemoryGiB"`
}

type SystemsLinks struct {
	Chassis   []base.Link
	ManagedBy []base.Link
}

type SystemsBoot struct {
}
