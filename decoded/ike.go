type Ike struct {
	V1 Ike_V1 `json:"v1"`
	V2 Ike_V2 `json:"v2"`
}

type Ike_V2 struct {
	AcceptedProposal   bool     `json:"accepted_proposal"`
	NotifyMessageTypes []uint64 `json:"notify_message_types"`
	VendorIds          []string `json:"vendor_ids"`
}

type Ike_V1 struct {
	AcceptedProposal   bool     `json:"accepted_proposal"`
	NotifyMessageTypes []uint64 `json:"notify_message_types"`
	VendorIds          []string `json:"vendor_ids"`
}
