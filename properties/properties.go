package properties

import "time"

// PropertyDetails yeah!
type PropertyDetails struct {
	ID                   uint `gorm:"primary_key",json:"uid"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time  `sql:"index"`
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
	AgencyID             int64       `json:"agencyId"`
	AgencyName           string      `json:"agencyName"`
	AgencyProfilePageURL string      `json:"agencyProfilePageUrl"`
	PropertyDetailsURL   string      `json:"propertyDetailsUrl"`
	DomainID int `json:"id"`
}

// GeoLocation yo!
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
