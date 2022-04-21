package decoded

type Snmp struct {
	OidInterfaces Snmp_OidInterfaces `json:"oid_interfaces"`
	OidPhysical   Snmp_OidPhysical   `json:"oid_physical"`
	OidSystem     Snmp_OidSystem     `json:"oid_system"`
}

type Snmp_OidSystem struct {
	Contact  string                  `json:"contact"`
	Desc     string                  `json:"desc"`
	InitTime uint64                  `json:"init_time"`
	Location string                  `json:"location"`
	Name     string                  `json:"name"`
	ObjectId string                  `json:"object_id"`
	Services Snmp_OidSystem_Services `json:"services"`
}

type Snmp_OidSystem_Services struct {
	Layer1 bool `json:"layer_1"`
	Layer2 bool `json:"layer_2"`
	Layer3 bool `json:"layer_3"`
	Layer4 bool `json:"layer_4"`
	Layer5 bool `json:"layer_5"`
	Layer6 bool `json:"layer_6"`
	Layer7 bool `json:"layer_7"`
}

type Snmp_OidPhysical struct {
	FirmwareRev string `json:"firmware_rev"`
	HardwareRev string `json:"hardware_rev"`
	MfgName     string `json:"mfg_name"`
	ModelName   string `json:"model_name"`
	Name        string `json:"name"`
	SerialNum   string `json:"serial_num"`
	SoftwareRev string `json:"software_rev"`
}

type Snmp_OidInterfaces struct {
	NumIfaces uint64 `json:"num_ifaces"`
}
