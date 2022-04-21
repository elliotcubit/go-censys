package decoded

type Rdp struct {
	CertificateInfo          Rdp_CertificateInfo          `json:"certificate_info"`
	ConnectResponse          Rdp_ConnectResponse          `json:"connect_response"`
	ProtocolFlags            Rdp_ProtocolFlags            `json:"protocol_flags"`
	SelectedSecurityProtocol Rdp_SelectedSecurityProtocol `json:"selected_security_protocol"`
	Version                  Rdp_Version                  `json:"version"`
	X224CcPduSrcref          uint64                       `json:"x224_cc_pdu_srcref"`
}

type Rdp_Version struct {
	Major int    `json:"major"`
	Minor int    `json:"minor"`
	Raw   uint64 `json:"raw"`
}

type Rdp_SelectedSecurityProtocol struct {
	Credssp                  bool   `json:"credssp"`
	CredsspEarlyAuth         bool   `json:"credssp_early_auth"`
	Error                    bool   `json:"error"`
	ErrorBadFlags            bool   `json:"error_bad_flags"`
	ErrorHybridRequired      bool   `json:"error_hybrid_required"`
	ErrorSslCertMissing      bool   `json:"error_ssl_cert_missing"`
	ErrorSslForbidden        bool   `json:"error_ssl_forbidden"`
	ErrorSslRequired         bool   `json:"error_ssl_required"`
	ErrorSslUserAuthRequired bool   `json:"error_ssl_user_auth_required"`
	ErrorUnknown             bool   `json:"error_unknown"`
	RawValue                 uint64 `json:"raw_value"`
	Rdstls                   bool   `json:"rdstls"`
	StandardRdp              bool   `json:"standard_rdp"`
	Tls                      bool   `json:"tls"`
}

type Rdp_ProtocolFlags struct {
	DynvcGraphicsPipeline       bool `json:"dynvc_graphics_pipeline"`
	ExtendedClientDataSupported bool `json:"extended_client_data_supported"`
	NegRespReserved             bool `json:"neg_resp_reserved"`
	RestrictedAdminMode         bool `json:"restricted_admin_mode"`
	RestrictedAuthMode          bool `json:"restricted_auth_mode"`
}

type Rdp_ConnectResponse struct {
	ConnectId        uint64                               `json:"connect_id"`
	DomainParameters Rdp_ConnectResponse_DomainParameters `json:"domain_parameters"`
}

type Rdp_ConnectResponse_DomainParameters struct {
	DomainProtocolVersion int64 `json:"domain_protocol_version"`
	MaxChannelIds         int64 `json:"max_channel_ids"`
	MaxMcspduSize         int64 `json:"max_mcspdu_size"`
	MaxProviderHeight     int64 `json:"max_provider_height"`
	MaxTokenIds           int64 `json:"max_token_ids"`
	MaxUserIdChannels     int64 `json:"max_user_id_channels"`
	MinThroughput         int64 `json:"min_throughput"`
	NumPriorities         int64 `json:"num_priorities"`
}

type Rdp_CertificateInfo struct {
	InternalX509ChainFps string                                `json:"internal_x509_chain_fps"`
	ProprietaryRsaKey    Rdp_CertificateInfo_ProprietaryRsaKey `json:"proprietary_rsa_key"`
}

type Rdp_CertificateInfo_ProprietaryRsaKey struct {
	KeyLength       uint64 `json:"key_length"`
	Magic           uint64 `json:"magic"`
	MaxBytesDatalen uint64 `json:"max_bytes_datalen"`
	Modulus         string `json:"modulus"`
	ModulusBitlen   uint64 `json:"modulus_bitlen"`
	PublicExponent  uint64 `json:"public_exponent"`
	Signature       string `json:"signature"`
}
