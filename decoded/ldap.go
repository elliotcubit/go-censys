package decoded

type Ldap struct {
	AllowsAnonymousBind bool             `json:"allows_anonymous_bind"`
	Attributes          []Ldap_Attribute `json:"attributes"`
	Resultcode          uint64           `json:"resultcode"`
}

type Ldap_Attribute struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
