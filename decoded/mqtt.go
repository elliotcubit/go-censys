package decoded

type Mqtt struct {
	ConnectionAckRaw      string                     `json:"connection_ack_raw"`
	ConnectionAckReturn   Mqtt_ConnectionAckReturn   `json:"connection_ack_return"`
	SubscriptionAckReturn Mqtt_SubscriptionAckReturn `json:"subscription_ack_return"`
}

type Mqtt_SubscriptionAckReturn struct {
	Raw         uint64 `json:"raw"`
	ReturnValue string `json:"return_value"`
}

type Mqtt_ConnectionAckReturn struct {
	Raw         uint64 `json:"raw"`
	ReturnValue string `json:"return_value"`
}
