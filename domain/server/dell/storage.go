package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

type StorageCollection struct {
	base.Meta
	Description       string `json:"Description"`
	Members           []base.Link
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`
}

type Storage struct {
	base.Meta
	Description                  string `json:"Description"`
	Drives                       []base.Link
	DrivesOdataCount             int    `json:"Drives@odata.count"`
	Id                           string `json:"Id"`
	Links                        SystemsStorageLinks
	Name                         string `json:"Name"`
	OEM                          SystemsStorageOEM
	Status                       base.Status
	StorageControllers           []SystemsStorageControllers
	StorageControllersOdataCount int `json:"StorageControllers@odata.count"`
	Volumes                      base.Link
}

func (s Storage) StatusToNumber() float64 {
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

type StorageDisk struct {
	base.Meta
	Actions                       SystemsStorageDiskAction
	Assembly                      base.Link
	BlockSizeBytes                int    `json:"BlockSizeBytes"`
	CapableSpeedGbs               int    `json:"CapableSpeedGbs"`
	CapacityBytes                 int    `json:"CapacityBytes"`
	Description                   string `json:"Description"`
	EncryptionAbility             string `json:"EncryptionAbility"`
	EncryptionStatus              string `json:"EncryptionStatus"`
	FailurePredicted              bool   `json:"FailurePredicted"`
	Id                            string `json:"Id"`
	Identifiers                   []SystemsStorageDiskIdentifiers
	IdentifiersOdataCount         int `json:"Identifiers@odata.count"`
	Links                         SystemsStorageDiskLinks
	Location                      []string `json:"Location"`
	Manufacturer                  string   `json:"Manufacturer"`
	MediaType                     string   `json:"MediaType"`
	Name                          string   `json:"Name"`
	NegotiatedSpeedGbs            int      `json:"NegotiatedSpeedGbs"`
	OEM                           SystemsStorageDiskOEM
	Operations                    []string `json:"Operations"`
	OperationsOdataCount          int      `json:"Operations@odata.count"`
	PartNumber                    string   `json:"PartNumber"`
	PhysicalLocation              SystemsStorageDiskPhysicalLocation
	PredictedMediaLifeLeftPercent float64 `json:"PredictedMediaLifeLeftPercent"`
	Protocol                      string  `json:"Protocol"`
	Revision                      string  `json:"Revision"`
	SerialNumber                  string  `json:"SerialNumber"`
	Status                        base.Status
	// RotationSpeedRPM              string  `json:"RotationSpeedRPM"`
}
