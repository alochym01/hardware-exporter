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
