package properties

// PropertyDetails yeah!
type PropertyDetails struct {
	UnitNumber           string      `json:"unitNumber"`
	StreetNumber         string      `json:"streetNumber"`
	StreetNamens         string      `json:"streetNamens"`
	StreetType           string      `json:"streetType"`
	Suburb               string      `json:"suburb"`
	Postcode             string      `json:"postcode"`
	State                string      `json:"state"`
	GeoLocation          GeoLocation `json:"geoLocation"`
	PropertyType         string      `json:"propertyType"`
	Bedrooms             int         `json:"bedrooms"`
	Bathrooms            int         `json:"bathrooms"`
	Carspaces            int         `json:"carspaces"`
	Price                int         `json:"price"`
	Result               string      `json:"result"`
	Agent                string      `json:"agent"`
	ID                   int64       `json:"id"`
	AgencyID             int64       `json:"agencyId"`
	AgencyName           string      `json:"agencyName"`
	AgencyProfilePageURL string      `json:"agencyProfilePageUrl"`
	PropertyDetailsURL   string      `json:"propertyDetailsUrl"`
}

// GeoLocation yo!
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
