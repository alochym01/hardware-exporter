package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

// Chassis Start
type ChasPowerControlPowerLimit struct {
	CorrectionInMs int    `json:"CorrectionInMs"`
	LimitException string `json:"LimitException"`
	LimitInWatts   int    `json:"LimitInWatts"`
}
type ChasPowerControlPowerMetrics struct {
	AverageConsumedWatts int `json:"AverageConsumedWatts"`
	IntervalInMin        int `json:"IntervalInMin"`
	MaxConsumedWatts     int `json:"MaxConsumedWatts"`
	MinConsumedWatts     int `json:"MinConsumedWatts"`
}
type ChasPowerControl struct {
	base.Meta
	MemberId            int `json:"MemberId"`
	Name                int `json:"Name"`
	PowerAllocatedWatts int `json:"PowerAllocatedWatts"`
	PowerAvailableWatts int `json:"PowerAvailableWatts"`
	PowerCapacityWatts  int `json:"PowerCapacityWatts"`
	PowerConsumedWatts  int `json:"PowerConsumedWatts"`
	PowerLimit          ChasPowerControlPowerLimit
	PowerMetrics        ChasPowerControlPowerMetrics
	PowerRequestedWatts int `json:"PowerRequestedWatts"`
	RelatedItem         []base.Link
}
type ChasPowerSuppliesInputRanges struct {
	InputType          string `json:"InputType"`
	MaximumFrequencyHz int    `json:"MaximumFrequencyHz"`
	MaximumVoltage     int    `json:"MaximumVoltage"`
	MinimumFrequencyHz int    `json:"MinimumFrequencyHz"`
	MinimumVoltage     int    `json:"MinimumVoltage"`
	OutputWattage      int    `json:"OutputWattage"`
}
type ChasPowerSuppliesOem struct {
}
type ChasPowerSupplies struct {
	base.Meta
	Assembly             base.Link
	EfficiencyPercent    float64 `json:"EfficiencyPercent"`
	FirmwareVersion      string  `json:"FirmwareVersion"`
	HotPluggable         bool    `json:"HotPluggable"`
	InputRanges          []ChasPowerSuppliesInputRanges
	LastPowerOutputWatts float64 `json:"LastPowerOutputWatts"`
	LineInputVoltage     float64 `json:"LineInputVoltage"`
	LineInputVoltageType string  `json:"LineInputVoltageType"`
	Manufacturer         string  `json:"Manufacturer"`
	MemberId             string  `json:"MemberId"`
	Model                string  `json:"Model"`
	Name                 string  `json:"MemberId"`
	OEM                  ChasPowerSuppliesOem
	PartNumber           string  `json:"PartNumber"`
	PowerCapacityWatts   float64 `json:"PowerCapacityWatts"`
	PowerInputWatts      float64 `json:"PowerInputWatts"`
	PowerOutputWatts     float64 `json:"PowerOutputWatts"`
	PowerSupplyType      string  `json:"PowerSupplyType"`
	Redundancy           []ChasPowerRedundancy
	RedundancyOdataCount int `json:"Redundancy@odata.count"`
	RelatedItem          []base.Link
	SerialNumber         string `json:"SerialNumber"`
	SparePartNumber      string `json:"SparePartNumber"`
	Status               base.Status
}
type ChasPowerRedundancy struct {
	base.Meta
	MaxNumSupported         int    `json:"MaxNumSupported"`
	MemberId                string `json:"MemberId"`
	MinNumNeeded            int    `json:"MinNumNeeded"`
	Mode                    string `json:"Mode"`
	Name                    string `json:"Name"`
	RedundancySet           []base.Link
	RedundancySetOdataCount int `json:"RedundancySet@odata.count"`
	Status                  base.Status
}
type ChasVoltages struct {
}

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

// Systems
type SystemTrustedModules struct {
	FirmwareVersion string `json:"FirmwareVersion"`
	InterfaceType   string `json:"InterfaceType"`
	Status          base.StateStatus
}

type SystemsProcessorSummary struct {
	Count                 int    `json:"Count"`
	LogicalProcessorCount int    `json:"LogicalProcessorCount"`
	Model                 string `json:"Model"`
	Status                base.Status
}

type SystemsOEM struct {
	// TODO take time to parse
	// {
	//     "Dell": {
	//         "DellSystem": {
	//             "@odata.context": "/redfish/v1/$metadata#DellSystem.DellSystem",
	//             "@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellSystem/System.Embedded.1",
	//             "@odata.type": "#DellSystem.v1_1_0.DellSystem",
	//             "BIOSReleaseDate": "01/18/2020",
	//             "BaseBoardChassisSlot": "NA",
	//             "BatteryRollupStatus": "OK",
	//             "BladeGeometry": "NotApplicable",
	//             "CMCIP": null,
	//             "CPURollupStatus": "OK",
	//             "ChassisModel": "",
	//             "ChassisName": "Main System Chassis",
	//             "ChassisServiceTag": "2GWL643",
	//             "ChassisSystemHeightUnit": 2,
	//             "CurrentRollupStatus": "OK",
	//             "EstimatedExhaustTemperatureCel": 43,
	//             "EstimatedSystemAirflowCFM": 32,
	//             "ExpressServiceCode": "5375758899",
	//             "FanRollupStatus": "OK",
	//             "IDSDMRollupStatus": null,
	//             "IntrusionRollupStatus": "OK",
	//             "IsOEMBranded": "False",
	//             "LastSystemInventoryTime": "2020-07-27T14:28:08+00:00",
	//             "LastUpdateTime": "2020-05-10T17:55:49+00:00",
	//             "LicensingRollupStatus": "OK",
	//             "MaxCPUSockets": 2,
	//             "MaxDIMMSlots": 16,
	//             "MaxPCIeSlots": 6,
	//             "MemoryOperationMode": "OptimizerMode",
	//             "NodeID": "2GWL643",
	//             "PSRollupStatus": "OK",
	//             "PopulatedDIMMSlots": 4,
	//             "PopulatedPCIeSlots": 1,
	//             "PowerCapEnabledState": "Disabled",
	//             "SDCardRollupStatus": null,
	//             "SELRollupStatus": "OK",
	//             "ServerAllocationWatts": null,
	//             "StorageRollupStatus": "OK",
	//             "SysMemErrorMethodology": "Multi-bitECC",
	//             "SysMemFailOverState": "NotInUse",
	//             "SysMemLocation": "SystemBoardOrMotherboard",
	//             "SysMemPrimaryStatus": "OK",
	//             "SystemGeneration": "14G Monolithic",
	//             "SystemID": 2242,
	//             "SystemRevision": "I",
	//             "TempRollupStatus": "OK",
	//             "TempStatisticsRollupStatus": "OK",
	//             "UUID": "4c4c4544-0047-5710-804c-b2c04f363433",
	//             "VoltRollupStatus": "OK",
	//             "smbiosGUID": "44454c4c-4700-1057-804c-b2c04f363433"
	//         }
	// 	}
	// }
}

type SystemsLinksOEM struct {
	// TODO take time to parse
	// {
	// 	"Dell": {
	// 		"BootOrder": {
	// 			"@odata.id": "/redfish/v1/Systems/System.Embedded.1/BootSources"
	// 		},
	// 		"DellBIOSService": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellBIOSService"
	// 		},
	// 		"DellChassisCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Chassis/System.Embedded.1/DellChassisCollection"
	// 		},
	// 		"DellGPUSensorCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellGPUSensorCollection"
	// 		},
	// 		"DellMetricService": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellMetricService"
	// 		},
	// 		"DellNumericSensorCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellNumericSensorCollection"
	// 		},
	// 		"DellOSDeploymentService": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellOSDeploymentService"
	// 		},
	// 		"DellPSNumericSensorCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellPSNumericSensorCollection"
	// 		},
	// 		"DellPresenceAndStatusSensorCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellPresenceAndStatusSensorCollection"
	// 		},
	// 		"DellRaidService": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellRaidService"
	// 		},
	// 		"DellSensorCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellSensorCollection"
	// 		},
	// 		"DellSlotCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellSlotCollection"
	// 		},
	// 		"DellSoftwareInstallationService": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellSoftwareInstallationService"
	// 		},
	// 		"DellVideoCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellVideoCollection"
	// 		},
	// 		"DellVideoNetworkCollection": {
	// 			"@odata.id": "/redfish/v1/Dell/Systems/System.Embedded.1/DellVideoNetworkCollection"
	// 		}
	// 	}
	// }

}
type SystemsLinks struct {
	Chassis             []base.Link
	ChassisOdataCount   int `json:"Chassis@odata.count"`
	CooledBy            []base.Link
	CooledByOdataCount  int         `json:"CooledBy@odata.count"`
	ManagedBy           []base.Link `json:"ManagedBy"`
	ManagedByOdataCount int         `json:"ManagedBy@odata.count"`
	PoweredBy           []base.Link
	PoweredByOdataCount int `json:"PoweredBy@odata.count"`
	OEM                 SystemsLinksOEM
}

type SystemHostWatchdogTimer struct {
	FunctionEnabled bool `json:"FunctionEnabled"`
	Status          base.StateStatus
	TimeoutAction   string `json:"TimeoutAction"`
}

type SystemsBoot struct {
	// TODO take time to parse
	// {
	//     "BootOptions": {
	//         "@odata.id": "/redfish/v1/Systems/System.Embedded.1/BootOptions"
	//     },
	//     "BootOrder": [
	//         "Boot0003",
	//         "Boot0000"
	//     ],
	//     "BootOrder@odata.count": 2,
	//     "BootSourceOverrideEnabled": "Disabled",
	//     "BootSourceOverrideMode": "UEFI",
	//     "BootSourceOverrideTarget": "None",
	//     "BootSourceOverrideTarget@Redfish.AllowableValues": [
	//         "None",
	//         "Pxe",
	//         "Floppy",
	//         "Cd",
	//         "Hdd",
	//         "BiosSetup",
	//         "Utilities",
	//         "UefiTarget",
	//         "SDCard",
	//         "UefiHttp"
	//     ],
	//     "UefiTargetBootSourceOverride": null
	// }
}

type SystemsMemorySummary struct {
	MemoryMirroring      string `json:"MemoryMirroring"`
	Status               base.Status
	TotalSystemMemoryGiB float64 `json:"TotalSystemMemoryGiB"`
}

type SysStorageCollectionMembers struct {
}

type SystemsStorageControllers struct {
}

type SystemsStorageOEM struct {
}

type SystemsStorageLinks struct {
}

// SystemsStorageDisk start
type SystemsStorageDiskAction struct {
}
type SystemsStorageDiskIdentifiers struct {
}
type SystemsStorageDiskLinks struct {
}
type SystemsStorageDiskOEM struct {
}
type SystemsStorageDiskPhysicalLocation struct {
	PartLocation SystemsStorageDiskPhysicalLocationPartLocation
}
type SystemsStorageDiskPhysicalLocationPartLocation struct {
	LocationOrdinalValue int    `json:"LocationOrdinalValue"`
	LocationType         string `json:"LocationType"`
}

// SystemsStorageDisk end

type SystemsEthernetInterfaceLinks struct {
}
