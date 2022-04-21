package censys

type TransportFingerprint struct {
	ID   int                      `json:"id"`
	OS   string                   `json:"os"`
	QUIC QuicTransportFingerprint `json:"quic"`
	Raw  string                   `json:"raw"`
}

type QuicTransportFingerprint struct {
	Versions []string `json:"versions"`
}
