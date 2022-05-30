package decoded

import "time"

type Jarm struct {
	CipherAndVersionFingerprint string    `json:"cipher_and_version_fingerprint"`
	Fingerprint                 string    `json:"fingerprins"`
	ObservedAt                  time.Time `json:"observed_at"`
	TlsExtensionsSHA256         string    `json:"tls_extensions_sha256"`
}
