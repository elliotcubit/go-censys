package decoded

type Ipp struct {
	AttributeCupsVersion string           `json:"attribute_cups_version"`
	AttributeIppVersions string           `json:"attribute_ipp_versions"`
	AttributePrinterUris string           `json:"attribute_printer_uris"`
	Attributes           []Ipp_Attribute  `json:"attributes"`
	CupsResponse         Ipp_CupsResponse `json:"cups_response"`
	CupsVersion          string           `json:"cups_version"`
	MajorVersion         uint64           `json:"major_version"`
	MinorVersion         uint64           `json:"minor_version"`
	Response             Ipp_Response     `json:"response"`
	VersionString        string           `json:"version_string"`
}

type Ipp_Response struct {
	Headers    HttpHeaders `json:"headers"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
}

type Ipp_CupsResponse struct {
	Headers    HttpHeaders `json:"headers"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
}

type Ipp_Attribute struct {
	Name     string `json:"name"`
	ValueTag uint64 `json:"value_tag"`
}
