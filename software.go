package censys

type Software struct {
	ComponentUniformResourceIdentifiers string               `json:"component_uniform_resource_identifiers"`
	Edition                             string               `json:"edition"`
	Language                            string               `json:"language"`
	Other                               map[string]string    `json:"other"`
	Part                                string               `json:"part"`
	Product                             string               `json:"product"`
	Source                              string               `json:"source"`
	SoftwareEdition                     string               `json:"sw_edition"`
	TargetHardware                      string               `json:"target_hw"`
	TargetSoftware                      string               `json:"target_sw"`
	TransportFingerprint                TransportFingerprint `json:"transport_fingerprint"`
	UniformResourceIdentifier           string               `json:"uniform_resource_identifier"`
	Update                              string               `json:"update"`
	Vendor                              string               `json:"vendor"`
	Version                             string               `json:"version"`
}
