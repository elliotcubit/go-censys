package decoded

type Pptp struct {
	BearerMessage   Pptp_Message  `json:"bearer_message"`
	ErrorMessage    Pptp_Message  `json:"error_message"`
	Firmware        Pptp_Firmware `json:"firmware"`
	FramingMessage  Pptp_Message  `json:"framing_message"`
	Hostname        string        `json:"hostname"`
	MaximumChannels uint64        `json:"maximum_channels"`
	Protocol        Pptp_Protocol `json:"protocol"`
	ResultMessage   Pptp_Message  `json:"result_message"`
	Vendor          string        `json:"vendor"`
}

type Pptp_Message struct {
	Code    uint64 `json:"code"`
	Meaning string `json:"meaning"`
}

type Pptp_Protocol struct {
	Major uint64 `json:"major"`
	Minor uint64 `json:"minor"`
}

type Pptp_Firmware struct {
	Major uint64 `json:"major"`
	Minor uint64 `json:"minor"`
}
