package decoded

type Amqp struct {
	ExplicitTLS bool            `json:"explicit_tls"`
	ImplicitTLS bool            `json:"implicit_tls"`
	ProtocolID  Amqp_ProtocolID `json:"protocol_id"`
	Version     Amqp_Version    `json:"version"`
}

type Amqp_ProtocolID struct {
	ID   uint64
	Name string
}

type Amqp_Version struct {
	Major    uint64 `json:"major"`
	Minor    uint64 `json:"minor"`
	Revision uint64 `json:"revision"`
}
