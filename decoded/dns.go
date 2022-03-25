package decoded

type DNS struct {
	ServerType        string         `json:"server_type"`
	ResolvesCorrectly bool           `json:"resolves_correctly"`
	Answers           []DNS_Answer   `json:"answers"`
	Questions         []DNS_Question `json:"questions"`
	EDNS              DNS_EDNS       `json:"edns"`
}

type DNS_Answer struct {
	Name     string `json:"name"`
	Response string `json:"response"`
	Type     string `json:"type"`
}

type DNS_Question struct {
	Name     string `json:"name"`
	Response string `json:"response"`
	Type     string `json:"type"`
}

type DNS_EDNS struct {
	Do      bool `json:"do"`
	UDP     int  `json:"udp"`
	Version int  `json:"version"`
}
