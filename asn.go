package censys

type AutonomousSystem struct {
	ASN         int    `json:"asn"`
	Description string `json:"description"`
	BGPPrefix   string `json:"bgp_prefix"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
}
