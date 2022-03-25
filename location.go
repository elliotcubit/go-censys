package censys

type Location struct {
	City                  string      `json:"city"`
	Continent             string      `json:"continent"`
	Country               string      `json:"country"`
	CountryCode           string      `json:"country_code"`
	PostalCode            string      `json:"postal_code"`
	Province              string      `json:"province"`
	Coordinates           Coordinates `json:"coordinates"`
	RegisteredCountry     string      `json:"registered_country"`
	RegisteredCountryCode string      `json:"registered_country_code"`
	Timezone              string      `json:"timezone"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
