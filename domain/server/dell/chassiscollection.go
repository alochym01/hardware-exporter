package dell

type ChassisCollection struct {
	ODataContext      string `json:"@odata.context"`
	ODataID           string `json:"@odata.id"`
	ODataType         string `json:"@odata.type"`
	Description       string `json:"Description"`
	Members           []Link
	MembersOdataCount int    `json:"Members@odata.count"`
	Name              string `json:"Name"`
}
