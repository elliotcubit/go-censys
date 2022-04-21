package decoded

type Sip struct {
	Code    int    `json:"code"`
	Server  string `json:"server"`
	Status  string `json:"status"`
	Version string `json:"version"`
}
