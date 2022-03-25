package decoded

type AnyConnect struct {
	AggregateAuthVersion int      `json:"aggregate_auth_version"`
	AuthMethods          []string `json:"auth_methods"`
	Groups               []string `json:"groups"`
	Raw                  string   `json:"raw"`
	ResponseType         string   `json:"response_type"`
}
