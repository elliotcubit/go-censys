package decoded

type Ssh struct {
	AlgorithmSelection Ssh_AlgorithmSelection `json:"algorithm_selection"`
	EndpointId         Ssh_EndpointId         `json:"endpoint_id"`
	KexInitMessage     Ssh_KexInitMessage     `json:"kex_init_message"`
	ServerHostKey      Ssh_ServerHostKey      `json:"server_host_key"`
}

type Ssh_ServerHostKey struct {
	CertkeyPublicKey  string                             `json:"certkey_public_key"`
	DsaPublicKey      Ssh_ServerHostKey_DsaPublicKey     `json:"dsa_public_key"`
	EcdsaPublicKey    Ssh_ServerHostKey_EcdsaPublicKey   `json:"ecdsa_public_key"`
	Ed25519PublicKey  Ssh_ServerHostKey_Ed25519PublicKey `json:"ed25519_public_key"`
	FingerprintSha256 string                             `json:"fingerprint_sha256"`
	RsaPublicKey      Ssh_ServerHostKey_RsaPublicKey     `json:"rsa_public_key"`
}

type Ssh_ServerHostKey_RsaPublicKey struct {
	Exponent string `json:"exponent"`
	Length   uint64 `json:"length"`
	Modulus  string `json:"modulus"`
}

type Ssh_ServerHostKey_Ed25519PublicKey struct {
	PublicBytes string `json:"public_bytes"`
}

type Ssh_ServerHostKey_EcdsaPublicKey struct {
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

type Ssh_ServerHostKey_DsaPublicKey struct {
	G string `json:"g"`
	P string `json:"p"`
	Q string `json:"q"`
	Y string `json:"y"`
}

type Ssh_KexInitMessage struct {
	ClientToServerCiphers     string `json:"client_to_server_ciphers"`
	ClientToServerCompression string `json:"client_to_server_compression"`
	ClientToServerLanguages   string `json:"client_to_server_languages"`
	ClientToServerMacs        string `json:"client_to_server_macs"`
	FirstKexFollows           bool   `json:"first_kex_follows"`
	HostKeyAlgorithms         string `json:"host_key_algorithms"`
	KexAlgorithms             string `json:"kex_algorithms"`
	ServerToClientCiphers     string `json:"server_to_client_ciphers"`
	ServerToClientCompression string `json:"server_to_client_compression"`
	ServerToClientLanguages   string `json:"server_to_client_languages"`
	ServerToClientMacs        string `json:"server_to_client_macs"`
}

type Ssh_EndpointId struct {
	Comment         string `json:"comment"`
	ProtocolVersion string `json:"protocol_version"`
	Raw             string `json:"raw"`
	SoftwareVersion string `json:"software_version"`
}

type Ssh_AlgorithmSelection struct {
	ClientToServerAlgGroup Ssh_AlgorithmSelection_ClientToServerAlgGroup `json:"client_to_server_alg_group"`
	HostKeyAlgorithm       string                                        `json:"host_key_algorithm"`
	KexAlgorithm           string                                        `json:"kex_algorithm"`
	ServerToClientAlgGroup Ssh_AlgorithmSelection_ServerToClientAlgGroup `json:"server_to_client_alg_group"`
}

type Ssh_AlgorithmSelection_ServerToClientAlgGroup struct {
	Cipher      string `json:"cipher"`
	Compression string `json:"compression"`
	Mac         string `json:"mac"`
}

type Ssh_AlgorithmSelection_ClientToServerAlgGroup struct {
	Cipher      string `json:"cipher"`
	Compression string `json:"compression"`
	Mac         string `json:"mac"`
}
