package censys

type TransportFingerprint struct {
	ID   int    `json:"id"`
	OS   string `json:"os"`
	QUIC string `json:"quic"`
	Raw  string `json:"raw"`
}

type QuicTransportFingerprint struct {
	Versions []string
}
