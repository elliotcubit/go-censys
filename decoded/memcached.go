package decoded

type Memcached struct {
	AsciiBindingProtocolEnabled  bool            `json:"ascii_binding_protocol_enabled"`
	BinaryBindingProtocolEnabled bool            `json:"binary_binding_protocol_enabled"`
	RespondsToUdp                bool            `json:"responds_to_udp"`
	Stats                        Memcached_Stats `json:"stats"`
	Version                      string          `json:"version"`
}

type Memcached_Stats struct {
	AuthEnabledSasl bool              `json:"auth_enabled_sasl"`
	Libevent        string            `json:"libevent"`
	Other           map[string]string `json:"other"`
}
