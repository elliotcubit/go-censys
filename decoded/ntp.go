package decoded

type Ntp struct {
	GetTimeHeader Ntp_GetTimeHeader `json:"get_time_header"`
}

type Ntp_GetTimeHeader struct {
	LeapIndicator uint64 `json:"leap_indicator"`
	Mode          uint64 `json:"mode"`
	Poll          int    `json:"poll"`
	Precision     int    `json:"precision"`
	ReferenceId   string `json:"reference_id"`
	Stratum       uint64 `json:"stratum"`
	Version       uint64 `json:"version"`
}
