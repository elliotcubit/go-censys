type Fox struct {
	AppName       string `json:"app_name"`
	AppVersion    string `json:"app_version"`
	AuthAgentType string `json:"auth_agent_type"`
	BrandId       string `json:"brand_id"`
	HostAddress   string `json:"host_address"`
	Hostid        string `json:"hostid"`
	Hostname      string `json:"hostname"`
	Id            uint64 `json:"id"`
	Language      string `json:"language"`
	OsName        string `json:"os_name"`
	OsVersion     string `json:"os_version"`
	StationName   string `json:"station_name"`
	SysInfo       string `json:"sys_info"`
	TimeZone      string `json:"time_zone"`
	Version       string `json:"version"`
	VmName        string `json:"vm_name"`
	VmUuid        string `json:"vm_uuid"`
	VmVersion     string `json:"vm_version"`
}
