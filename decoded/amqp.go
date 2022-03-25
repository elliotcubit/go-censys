package decoded

type AMQP struct {
	ExplicitTLS bool            `json:"explicit_tls"`
	ImplicitTLS bool            `json:"implicit_tls"`
	ProtocolID  AMQP_ProtocolID `json:"protocol_id"`
	Version     AMQP_Version    `json:"version"`
}

type AMQP_ProtocolID struct {
	ID   uint64
	Name string
}

type AMQP_Version struct {
	Major    uint64 `json:"major"`
	Minor    uint64 `json:"minor"`
	Revision uint64 `json:"revision"`
}
