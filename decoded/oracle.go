package decoded

type Oracle struct {
	AcceptVersion        uint64                      `json:"accept_version"`
	ConnectFlags0        Oracle_ConnectFlags         `json:"connect_flags0"`
	ConnectFlags1        Oracle_ConnectFlags         `json:"connect_flags1"`
	DidResend            bool                        `json:"did_resend"`
	GlobalServiceOptions Oracle_GlobalServiceOptions `json:"global_service_options"`
	NsnServiceVersions   Oracle_NsnServiceVersions   `json:"nsn_service_versions"`
	NsnVersion           string                      `json:"nsn_version"`
	RedirectTarget       []Oracle_RedirectTarget     `json:"redirect_target"`
	RedirectTargetRaw    string                      `json:"redirect_target_raw"`
	RefuseError          []Oracle_RefuseError        `json:"refuse_error"`
	RefuseErrorRaw       string                      `json:"refuse_error_raw"`
	RefuseReasonApp      string                      `json:"refuse_reason_app"`
	RefuseReasonSys      string                      `json:"refuse_reason_sys"`
	RefuseVersion        string                      `json:"refuse_version"`
}

type Oracle_RefuseError struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Oracle_RedirectTarget struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Oracle_NsnServiceVersions struct {
	Authentication string `json:"authentication"`
	DataIntegrity  string `json:"data_integrity"`
	Encryption     string `json:"encryption"`
	Supervisor     string `json:"supervisor"`
}

type Oracle_GlobalServiceOptions struct {
	AttentionProcessing bool `json:"attention_processing"`
	BrokenConnectNotify bool `json:"broken_connect_notify"`
	CanReceiveAttention bool `json:"can_receive_attention"`
	CanSendAttention    bool `json:"can_send_attention"`
	DirectIo            bool `json:"direct_io"`
	FullDuplex          bool `json:"full_duplex"`
	HalfDuplex          bool `json:"half_duplex"`
	HeaderChecksum      bool `json:"header_checksum"`
	PacketChecksum      bool `json:"packet_checksum"`
	Unknown0001         bool `json:"unknown_0001"`
	Unknown0020         bool `json:"unknown_0020"`
	Unknown0040         bool `json:"unknown_0040"`
	Unknown0080         bool `json:"unknown_0080"`
	Unknown0100         bool `json:"unknown_0100"`
}

type Oracle_ConnectFlags struct {
	InterchangeInvolved bool `json:"interchange_involved"`
	ServicesEnabled     bool `json:"services_enabled"`
	ServicesLinkedIn    bool `json:"services_linked_in"`
	ServicesRequired    bool `json:"services_required"`
	ServicesWanted      bool `json:"services_wanted"`
	Unknown20           bool `json:"unknown_20"`
	Unknown40           bool `json:"unknown_40"`
	Unknown80           bool `json:"unknown_80"`
}
