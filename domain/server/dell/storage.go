package dell

import "github.com/alochym01/hardware-exporter/domain/server/base"

type StorageCollection struct {
	base.Meta
	Description       string `json:"Description"`
	Members           []base.Link
	MembersOdataCount string `json:"Members@odata.count"`
	Name              string `json:"Name"`
}
