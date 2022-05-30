package decoded

type TLS struct {
	Certificate       TLS_Certificate       `json:"certificates"`
	CipherSelected    string                `json:"cipher_selected"`
	Ja3s              string                `json:"ja3s"`
	ServerKeyExchange TLS_ServerKeyExchange `json:"server_key_exchange"`
	SessionTicket     TLS_SessionTicket     `json:"session_ticket"`
	VersionSelected   string                `json:"version_selected"`
}

type TLS_SessionTicket struct {
	Length       uint64 `json:"length"`
	LifetimeHint uint64 `json:"lifetime_hint"`
}

type TLS_ServerKeyExchange struct {
	DhParams  TLS_ServerKeyExchange_DhParams  `json:"dh_params"`
	EcParams  TLS_ServerKeyExchange_EcParams  `json:"ec_params"`
	RsaParams TLS_ServerKeyExchange_RsaParams `json:"rsa_params"`
	Signature string                          `json:"signature"`
}

type TLS_ServerKeyExchange_RsaParams struct {
	PublicKey TLS_ServerKeyExchange_RsaParams_PublicKey `json:"public_key"`
}

type TLS_ServerKeyExchange_RsaParams_PublicKey struct {
	E string `json:"e"`
	N string `json:"n"`
}

type TLS_ServerKeyExchange_EcParams struct {
	NamedCurve uint64 `json:"named_curve"`
	PublicKey  string `json:"public_key"`
}

type TLS_ServerKeyExchange_DhParams struct {
	Group     TLS_ServerKeyExchange_DhParams_Group `json:"group"`
	PublicKey string                               `json:"public_key"`
}

type TLS_ServerKeyExchange_DhParams_Group struct {
	P string `json:"p"`
}

type TLS_Certificate struct {
	Chain          []TLS_Certificate_Chain  `json:"chain"`
	ChainFpsSha256 []string                 `json:"chain_fps_sha_256"`
	LeafData       TLS_Certificate_LeafData `json:"leaf_data"`
	LeafFpSha256   string                   `json:"leaf_fp_sha_256"`
}

type TLS_Certificate_LeafData struct {
	Fingerprint     string                             `json:"fingerprint"`
	Issuer          TLS_Certificate_LeafData_Name      `json:"issuer"`
	IssuerDn        string                             `json:"issuer_dn"`
	Names           []string                           `json:"names"`
	PubkeyAlgorithm string                             `json:"pubkey_algorithm"`
	PubkeyBitSize   int                                `json:"pubkey_bit_size"`
	PublicKey       TLS_Certificate_LeafData_PublicKey `json:"public_key"`
	Signature       TLS_Certificate_LeafData_Signature `json:"signature"`
	Subject         TLS_Certificate_LeafData_Name      `json:"subject"`
	SubjectDn       string                             `json:"subject_dn"`
	TbsFingerprint  string                             `json:"tbs_fingerprint"`
}

type TLS_Certificate_LeafData_Signature struct {
	SelfSigned         bool   `json:"self_signed"`
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type TLS_Certificate_LeafData_PublicKey struct {
	Dsa          TLS_Certificate_LeafData_PublicKey_Dsa   `json:"dsa"`
	Ecdsa        TLS_Certificate_LeafData_PublicKey_Ecdsa `json:"ecdsa"`
	Fingerprint  string                                   `json:"fingerprint"`
	KeyAlgorithm string                                   `json:"key_algorithm"`
	Rsa          TLS_Certificate_LeafData_PublicKey_Rsa   `json:"rsa"`
}

type TLS_Certificate_LeafData_PublicKey_Rsa struct {
	Exponent string `json:"exponent"`
	Length   uint64 `json:"length"`
	Modulus  string `json:"modulus"`
}

type TLS_Certificate_LeafData_PublicKey_Ecdsa struct {
	B      string `json:"b"`
	Curve  string `json:"curve"`
	Gx     string `json:"gx"`
	Gy     string `json:"gy"`
	Length uint64 `json:"length"`
	N      string `json:"n"`
	P      string `json:"p"`
	Pub    string `json:"pub"`
	X      string `json:"x"`
	Y      string `json:"y"`
}

type TLS_Certificate_LeafData_PublicKey_Dsa struct {
	G string `json:"g"`
	P string `json:"p"`
	Q string `json:"q"`
	Y string `json:"y"`
}

type TLS_Certificate_LeafData_Name struct {
	CommonName           []string `json:"common_name"`
	Country              []string `json:"country"`
	DomainComponent      []string `json:"domain_component"`
	EmailAddress         []string `json:"email_address"`
	JurisdictionCountry  []string `json:"jurisdiction_country"`
	JurisdictionLocality []string `json:"jurisdiction_locality"`
	JurisdictionProvince []string `json:"jurisdiction_province"`
	Locality             []string `json:"locality"`
	Organization         []string `json:"organization"`
	OrganizationId       []string `json:"organization_id"`
	OrganizationalUnit   []string `json:"organizational_unit"`
	PostalCode           []string `json:"postal_code"`
	Province             []string `json:"province"`
	SerialNumber         []string `json:"serial_number"`
	StreetAddress        []string `json:"street_address"`
}

type TLS_Certificate_Chain struct {
	Fingerprint string `json:"fingerprint"`
	IssuerDn    string `json:"issuer_dn"`
	SubjectDn   string `json:"subject_dn"`
}
