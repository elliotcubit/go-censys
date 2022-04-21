package decoded

type S7 struct {
	Copyright          string `json:"copyright"`
	CpuProfile         string `json:"cpu_profile"`
	Firmware           string `json:"firmware"`
	Hardware           string `json:"hardware"`
	Location           string `json:"location"`
	MemorySerialNumber string `json:"memory_serial_number"`
	Module             string `json:"module"`
	ModuleId           string `json:"module_id"`
	ModuleType         string `json:"module_type"`
	OemId              string `json:"oem_id"`
	PlantId            string `json:"plant_id"`
	ReservedForOs      string `json:"reserved_for_os"`
	SerialNumber       string `json:"serial_number"`
	System             string `json:"system"`
}
