package decoded

type Upnp struct {
	Devices  []Upnp_Device `json:"devices"`
	Endpoint string        `json:"endpoint"`
	Headers  HttpHeaders   `json:"headers"`
	Spec     Upnp_Spec     `json:"spec"`
}

type Upnp_Spec struct {
	Major string `json:"major"`
	Minor string `json:"minor"`
}

type Upnp_Device struct {
	DeviceType       string                     `json:"device_type"`
	FriendlyName     string                     `json:"friendly_name"`
	Id               int                        `json:"id"`
	Manufacturer     string                     `json:"manufacturer"`
	ManufacturerUrl  string                     `json:"manufacturer_url"`
	ModelDescription string                     `json:"model_description"`
	ModelName        string                     `json:"model_name"`
	ModelNumber      string                     `json:"model_number"`
	ModelUrl         string                     `json:"model_url"`
	ParentId         int                        `json:"parent_id"`
	PresentationUrl  string                     `json:"presentation_url"`
	SerialNumber     string                     `json:"serial_number"`
	ServiceList      []Upnp_Devices_ServiceList `json:"service_list"`
	Udn              string                     `json:"udn"`
	Upc              string                     `json:"upc"`
}

type Upnp_Devices_ServiceList struct {
	ControlUrl  string `json:"control_url"`
	EventSubUrl string `json:"event_sub_url"`
	ScpdUrl     string `json:"scpd_url"`
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
}
