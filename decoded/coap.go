package decoded

type Coap struct {
	Code        string `json:"code"`
	MessageID   uint64 `json:"message_id"`
	MessageType string `json:"message_type"`
	Payload     string `json:"payload"`
	Token       string `json:"token"`
	Version     string `json:"version"`
}
