package decoded

type Dns struct {
	ServerType        string         `json:"server_type"`
	ResolvesCorrectly bool           `json:"resolves_correctly"`
	Answers           []Dns_Answer   `json:"answers"`
	Questions         []Dns_Question `json:"questions"`
	EDNS              Dns_EDNS       `json:"edns"`
}

type Dns_Answer struct {
	Name     string `json:"name"`
	Response string `json:"response"`
	Type     string `json:"type"`
}

type Dns_Question struct {
	Name     string `json:"name"`
	Response string `json:"response"`
	Type     string `json:"type"`
}

type Dns_EDNS struct {
	Do      bool `json:"do"`
	UDP     int  `json:"udp"`
	Version int  `json:"version"`
}
