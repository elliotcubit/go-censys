package decoded

type Bacnet struct {
	ApplicationSoftwareRevision string `json:"application_software_revision"`
	Description                 string `json:"description"`
	FirmwareRevision            string `json:"firmware_revision"`
	InstanceNumber              uint64 `json:"instance_number"`
	Location                    string `json:"location"`
	ModelName                   string `json:"model_name"`
	ObjectName                  string `json:"object_name"`
	VendorID                    uint64 `json:"vendor_id"`
	VendorName                  string `json:"vendor_name"`
}
