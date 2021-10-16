package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

type ChassisThermal struct {
	base.Meta
	Fans                   []ChassisThermalFans
	FansOdataCount         string `json:"Fans@odata.count"`
	Id                     string `json:"Id"`
	Name                   string `json:"Name"`
	Redundancy             []ChassisThermalRedundancy
	RedundancyOdataCount   string `json:"Redundancy@odata.count"`
	Temperatures           []ChassisThermalTemperatures
	TemperaturesOdataCount string `json:"Temperatures@odata.count"`
}
