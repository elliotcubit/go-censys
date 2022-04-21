package decoded

type Ipmi struct {
	Capabilities   Ipmi_Capabilities   `json:"capabilities"`
	CommandPayload Ipmi_CommandPayload `json:"command_payload"`
	Raw            string              `json:"raw"`
	RmcpHeader     Ipmi_RmcpHeader     `json:"rmcp_header"`
	SessionHeader  Ipmi_SessionHeader  `json:"session_header"`
}

type Ipmi_SessionHeader struct {
	AuthCode              string                      `json:"auth_code"`
	AuthType              Ipmi_SessionHeader_AuthType `json:"auth_type"`
	SessionId             int64                       `json:"session_id"`
	SessionSequenceNumber int64                       `json:"session_sequence_number"`
}

type Ipmi_SessionHeader_AuthType struct {
	Name string `json:"name"`
	Raw  int    `json:"raw"`
}

type Ipmi_RmcpHeader struct {
	MessageClass   Ipmi_RmcpHeader_MessageClass `json:"message_class"`
	SequenceNumber int                          `json:"sequence_number"`
	Version        int                          `json:"version"`
}

type Ipmi_RmcpHeader_MessageClass struct {
	Class int    `json:"class"`
	IsAck bool   `json:"is_ack"`
	Name  string `json:"name"`
	Raw   int    `json:"raw"`
}

type Ipmi_CommandPayload struct {
	ChecksumError           bool                                    `json:"checksum_error"`
	Data                    string                                  `json:"data"`
	IpmiCommandNumber       Ipmi_CommandPayload_IpmiCommandNumber   `json:"ipmi_command_number"`
	NetworkFunctionCode     Ipmi_CommandPayload_NetworkFunctionCode `json:"network_function_code"`
	RequestorSequenceNumber int                                     `json:"requestor_sequence_number"`
}

type Ipmi_CommandPayload_NetworkFunctionCode struct {
	Raw               int                                                       `json:"raw"`
	NetFn             Ipmi_CommandPayload_NetworkFunctionCode_NetFn             `json:"net_fn"`
	LogicalUnitNumber Ipmi_CommandPayload_NetworkFunctionCode_LogicalUnitNumber `json:"logical_unit_number"`
}

type Ipmi_CommandPayload_NetworkFunctionCode_NetFn struct {
	IsRequest  bool   `json:"is_request"`
	IsResponse bool   `json:"is_response"`
	Name       string `json:"name"`
	Raw        int    `json:"raw"`
	Value      int    `json:"value"`
}

type Ipmi_CommandPayload_NetworkFunctionCode_LogicalUnitNumber struct {
	Name string `json:"name"`
	Raw  int    `json:"raw"`
}

type Ipmi_CommandPayload_IpmiCommandNumber struct {
	Name string `json:"name"`
	Raw  int    `json:"raw"`
}

type Ipmi_Capabilities struct {
	AuthStatus           Ipmi_Capabilities_AuthStatus           `json:"auth_status"`
	ChannelNumber        int                                    `json:"channel_number"`
	CompletionCode       Ipmi_Capabilities_CompletionCode       `json:"completion_code"`
	ExtendedCapabilities Ipmi_Capabilities_ExtendedCapabilities `json:"extended_capabilities"`
	OemData              int                                    `json:"oem_data"`
	OemId                string                                 `json:"oem_id"`
	SupportedAuthTypes   Ipmi_Capabilities_SupportedAuthTypes   `json:"supported_auth_types"`
}

type Ipmi_Capabilities_SupportedAuthTypes struct {
	Extended       bool `json:"extended"`
	Md2            bool `json:"md2"`
	Md5            bool `json:"md5"`
	None           bool `json:"none"`
	OemProprietary bool `json:"oem_proprietary"`
	Password       bool `json:"password"`
	Raw            int  `json:"raw"`
}

type Ipmi_Capabilities_ExtendedCapabilities struct {
	SupportsIpmiV15 bool `json:"supports_ipmi_v1_5"`
	SupportsIpmiV20 bool `json:"supports_ipmi_v2_0"`
}

type Ipmi_Capabilities_CompletionCode struct {
	Name string `json:"name"`
	Raw  int    `json:"raw"`
}

type Ipmi_Capabilities_AuthStatus struct {
	AnonymousLoginEnabled bool `json:"anonymous_login_enabled"`
	AuthEachMessage       bool `json:"auth_each_message"`
	HasAnonymousUsers     bool `json:"has_anonymous_users"`
	HasNamedUsers         bool `json:"has_named_users"`
	TwoKeyLoginRequired   bool `json:"two_key_login_required"`
	UserAuthDisabled      bool `json:"user_auth_disabled"`
}
