package hpe

import "github.com/alochym01/hardware-exporter/domain/server/base"

type DiskDrivesFirmwareVersion struct {
}
type StorageDisk struct {
	base.Meta
	BlockSizeBytes                    int      `json:"BlockSizeBytes"`
	CapacityGB                        float64  `json:"CapacityGB"`
	CapacityLogicalBlocks             int      `json:"CapacityLogicalBlocks"`
	CapacityMiB                       float64  `json:"CapacityMiB"`
	CarrierApplicationVersion         string   `json:"CarrierApplicationVersion"`
	CarrierAuthenticationStatus       string   `json:"CarrierAuthenticationStatus"`
	CurrentTemperatureCelsius         int      `json:"CurrentTemperatureCelsius"`
	Description                       string   `json:"Description"`
	DiskDriveStatusReasons            []string `json:"DiskDriveStatusReasons"`
	DiskDriveUse                      string   `json:"DiskDriveUse"`
	EncryptedDrive                    bool     `json:"EncryptedDrive"`
	FirmwareVersion                   DiskDrivesFirmwareVersion
	Id                                string  `json:"Id"`
	InterfaceSpeedMbps                int     `json:"InterfaceSpeedMbps"`
	InterfaceType                     string  `json:"InterfaceType"`
	LegacyBootPriority                string  `json:"LegacyBootPriority"`
	Location                          string  `json:"Location"`
	LocationFormat                    string  `json:"LocationFormat"`
	MaximumTemperatureCelsius         int     `json:"MaximumTemperatureCelsius"`
	MediaType                         string  `json:"MediaType"`
	Model                             string  `json:"Model"`
	Name                              string  `json:"Name"`
	PowerOnHours                      int     `json:"PowerOnHours"`
	RotationalSpeedRpm                int     `json:"RotationalSpeedRpm"`
	SSDEnduranceUtilizationPercentage float64 `json:"SSDEnduranceUtilizationPercentage"`
	SerialNumber                      string  `json:"SerialNumber"`
	Status                            base.Status
	UncorrectedReadErrors             int `json:"UncorrectedReadErrors"`
	UncorrectedWriteErrors            int `json:"UncorrectedWriteErrors"`

	/*

		    "@odata.context": "/redfish/v1/$metadata#HpeSmartStorageDiskDrive.HpeSmartStorageDiskDrive",
		    "@odata.etag": "W/\"DF577A00\"",
		    "@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/DiskDrives/0",
		    "@odata.type": "#HpeSmartStorageDiskDrive.v2_1_0.HpeSmartStorageDiskDrive",
		    "BlockSizeBytes": 512,
		    "CapacityGB": 600,
		    "CapacityLogicalBlocks": 1172123568,
		    "CapacityMiB": 572325,
		    "CarrierApplicationVersion": "11",
		    "CarrierAuthenticationStatus": "OK",
		    "CurrentTemperatureCelsius": 38,
		    "Description": "HPE Smart Storage Disk Drive View",
		    "DiskDriveStatusReasons": [
		        "None"
		    ],
		    "DiskDriveUse": "Data",
		    "EncryptedDrive": false,
		    "FirmwareVersion": {
		        "Current": {
		            "VersionString": "HPD1"
		        }
		    },
		    "Id": "0",
		    "InterfaceSpeedMbps": 12000,
		    "InterfaceType": "SAS",
		    "LegacyBootPriority": "Primary",
		    "Location": "1I:1:1",
		    "LocationFormat": "ControllerPort:Box:Bay",
		    "MaximumTemperatureCelsius": 48,
		    "MediaType": "HDD",
		    "Model": "EG000600JWJNH",
		    "Name": "HpeSmartStorageDiskDrive",
		    "PowerOnHours": null,
		    "RotationalSpeedRpm": 10000,
		    "SSDEnduranceUtilizationPercentage": null,
		    "SerialNumber": "39Q0A0WDFF5F",
		    "Status": {
		        "Health": "OK",
		        "State": "Enabled"
		    },
		    "UncorrectedReadErrors": 0,
		    "UncorrectedWriteErrors": 0
		}

	*/
}

type SmartStorageDiskDriveCollection struct {
	StorageArrayControllerCollection
}
type StorageArrayControllerLinks struct {
	LogicalDrives      base.Link
	PhysicalDrives     base.Link
	StorageEnclosures  base.Link
	UnconfiguredDrives base.Link
}
type StorageArrayController struct {
	base.Meta
	Links  StorageArrayControllerLinks
	Status base.Status
	// {
	// 	"@odata.context": "/redfish/v1/$metadata#HpeSmartStorageArrayController.HpeSmartStorageArrayController",
	// 	"@odata.etag": "W/\"B29953BC\"",
	// 	"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/",
	// 	"@odata.type": "#HpeSmartStorageArrayController.v2_2_0.HpeSmartStorageArrayController",
	// 	"AdapterType": "SmartArray",
	// 	"BackupPowerSourceStatus": "Present",
	// 	"CacheMemorySizeMiB": 2048,
	// 	"CacheModuleSerialNumber": "               ",
	// 	"CacheModuleStatus": {
	// 		"Health": "OK"
	// 	},
	// 	"ControllerBoard": {
	// 		"Status": {
	// 			"Health": "OK"
	// 		}
	// 	},
	// 	"ControllerPartNumber": "836260-001",
	// 	"CurrentOperatingMode": "Mixed",
	// 	"Description": "HPE Smart Storage Array Controller View",
	// 	"DriveWriteCache": "Disabled",
	// 	"EncryptionCryptoOfficerPasswordSet": false,
	// 	"EncryptionCspTestPassed": true,
	// 	"EncryptionEnabled": false,
	// 	"EncryptionFwLocked": false,
	// 	"EncryptionHasLockedVolumesMissingBootPassword": false,
	// 	"EncryptionMixedVolumesEnabled": false,
	// 	"EncryptionSelfTestPassed": true,
	// 	"EncryptionStandaloneModeEnabled": false,
	// 	"ExternalPortCount": 0,
	// 	"FirmwareVersion": {
	// 		"Current": {
	// 			"VersionString": "1.98"
	// 		}
	// 	},
	// 	"HardwareRevision": "B",
	// 	"Id": "0",
	// 	"InternalPortCount": 2,
	// 	"Links": {
	// 		"LogicalDrives": {
	// 			"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/LogicalDrives/"
	// 		},
	// 		"PhysicalDrives": {
	// 			"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/DiskDrives/"
	// 		},
	// 		"StorageEnclosures": {
	// 			"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/StorageEnclosures/"
	// 		},
	// 		"UnconfiguredDrives": {
	// 			"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/UnconfiguredDrives/"
	// 		}
	// 	},
	// 	"Location": "Slot 0",
	// 	"LocationFormat": "PCISlot",
	// 	"Model": "HPE Smart Array P408i-a SR Gen10",
	// 	"Name": "HpeSmartStorageArrayController",
	// 	"ReadCachePercent": 10,
	// 	"SerialNumber": "PEYHC0DRHC73AL ",
	// 	"Status": {
	// 		"Health": "OK",
	// 		"State": "Enabled"
	// 	},
	// 	"WriteCacheBypassThresholdKB": 1040,
	// 	"WriteCacheWithoutBackupPowerEnabled": false
	// }

}
type StorageArrayControllerCollection struct {
	base.Meta
	ODataEtag         string `json:"@odata.etag"`
	Description       string `json:"Description"`
	Members           []base.Link
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`

	// {
	// 	"@odata.context": "/redfish/v1/$metadata#HpeSmartStorageArrayControllerCollection.HpeSmartStorageArrayControllerCollection",
	// 	"@odata.etag": "W/\"AA6D42B0\"",
	// 	"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/",
	// 	"@odata.type": "#HpeSmartStorageArrayControllerCollection.HpeSmartStorageArrayControllerCollection",
	// 	"Description": "HPE Smart Storage Array Controllers View",
	// 	"Members": [
	// 		{
	// 			"@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/0/"
	// 		}
	// 	],
	// 	"Members@odata.count": 1,
	// 	"Name": "HpeSmartStorageArrayControllers"
	// }

}
type StorageCollectionLinks struct {
	ArrayControllers base.Link
	HostBusAdapters  base.Link
}
type StorageCollection struct {
	base.Meta
	ODataEtag   string `json:"@odata.etag"`
	Id          string `json:"Id"`
	Links       StorageCollectionLinks
	Description string `json:"Description"`
	Members     []base.Link
	Status      base.HealthStatus
	Name        string `json:"Name"`

	// "@odata.context": "/redfish/v1/$metadata#HpeSmartStorage.HpeSmartStorage",
	// "@odata.etag": "W/\"BEFCD4EA\"",
	// "@odata.id": "/redfish/v1/Systems/1/SmartStorage/",
	// "@odata.type": "#HpeSmartStorage.v2_0_0.HpeSmartStorage",
	// "Description": "HPE Smart Storage",
	// "Id": "SmartStorage",
	// "Links": {
	//     "ArrayControllers": {
	//         "@odata.id": "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/"
	//     },
	//     "HostBusAdapters": {
	//         "@odata.id": "/redfish/v1/Systems/1/SmartStorage/HostBusAdapters/"
	//     }
	// },
	// "Name": "HpeSmartStorage",
	// "Status": {
	//     "Health": "OK"
	// }

}
type Storage struct {
	base.Meta
	Description      string `json:"Description"`
	Drives           []base.Link
	DrivesOdataCount int    `json:"Drives@odata.count"`
	Id               string `json:"Id"`
	// Links                        SystemsStorageLinks
	Name string `json:"Name"`
	// OEM                          SystemsStorageOEM
	Status base.Status
	// StorageControllers           []SystemsStorageControllers
	StorageControllersOdataCount int `json:"StorageControllers@odata.count"`
	Volumes                      base.Link
}

func (s StorageCollection) StatusToNumber() float64 {
	switch s.Status.Health {
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
