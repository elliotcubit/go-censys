package decoded

type Ftp struct {
	AuthSslResponse string `json:"auth_ssl_response"`
	AuthTlsResponse string `json:"auth_tls_response"`
	Banner          string `json:"banner"`
	ImplicitTls     bool   `json:"implicit_tls"`
	StatusCode      int    `json:"status_code"`
	StatusMeaning   string `json:"status_meaning"`
}
