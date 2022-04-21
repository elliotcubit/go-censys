package decoded

type PcAnywhere struct {
	Name   string            `json:"name"`
	Nr     string            `json:"nr"`
	Status PcAnywhere_Status `json:"status"`
}

type PcAnywhere_Status struct {
	InUse bool   `json:"in_use"`
	Raw   string `json:"raw"`
}
