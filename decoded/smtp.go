package decoded

type Smtp struct {
	Banner   string `json:"banner"`
	Ehlo     string `json:"ehlo"`
	StartTls string `json:"start_tls"`
}
