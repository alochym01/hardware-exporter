package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

type PowerControl struct {
	base.Meta
	Description             string                `json:"Description"`
	Id                      string                `json:"Id"`
	Name                    string                `json:"Name"`
	PowerControl            []ChasPowerControl    // Chassis
	PowerControlOdataCount  int                   `json:"PowerControl@odata.count"`
	PowerSupplies           []ChasPowerSupplies   // Chassis
	PowerSuppliesOdataCount int                   `json:"PowerSupplies@odata.count"`
	Redundancy              []ChasPowerRedundancy // Chassis
	RedundancyOdataCount    int                   `json:"Redundancy@odata.count"`
	Voltages                []ChasVoltages        // Chassis
	VoltagesOdataCount      int                   `json:"Voltages@odata.count"`
}
