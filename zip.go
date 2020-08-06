package geocodio

// Zip4 based on this payload example
/*
{
	"record_type": {
		"code": "S",
		"description": "Street"
	},
	"carrier_route": {
		"id": "C007",
		"description": "City Delivery"
	},
	"building_or_firm_name": null,
	"plus4": [
		"2890"
	],
	"zip9": [
		"22201-2890"
	],
	"government_building": null,
	"facility_code": {
		"code": "P",
		"description": "Post Office"
	},
	"city_delivery": true,
	"valid_delivery_area": true,
	"exact_match": true
}
*/

type Zip4 struct {
	RecodeType         RecordType   `json:"record_type,omitempty"`
	CarrierRoute       CarrierRoute `json:"carrier_route,omitempty"`
	BuildingOrFirmName string       `json:"building_or_firm_name,omitempty"`
	Plus4              []string     `json:"plus4,omitempty"`
	Zip9               []string     `json:"zip9,omitempty"`
	GovermentBuilding  string       `json:"government_building,omitempty"`
	FacilityCode       FacilityCode `json:"facility_code,omitempty"`
	CityDelivery       bool         `json:"city_delivery,omitempty"`
	ValidDeliveryArea  bool         `json:"valid_delivery_area,omitempty"`
	ExactMatch         bool         `json:"exact_match,omitempty"`
}

type RecordType struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type CarrierRoute struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type FacilityCode struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
