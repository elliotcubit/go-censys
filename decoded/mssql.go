package decoded

type Mssql struct {
	EncryptMode     string               `json:"encrypt_mode"`
	InstanceName    string               `json:"instance_name"`
	PreloginOptions Mssql_PreloginOption `json:"prelogin_options"`
	Version         string               `json:"version"`
}

type Mssql_PreloginOption struct {
	EncryptMode     string                              `json:"encrypt_mode"`
	FedAuthRequired bool                                `json:"fed_auth_required"`
	Instance        string                              `json:"instance"`
	Mars            bool                                `json:"mars"`
	Nonce           string                              `json:"nonce"`
	ServerVersion   Mssql_PreloginOptions_ServerVersion `json:"server_version"`
	ThreadId        uint64                              `json:"thread_id"`
	TraceId         string                              `json:"trace_id"`
	Unknown         map[uint64]string                   `json:"unknown"`
}

type Mssql_PreloginOptions_ServerVersion struct {
	BuildNumber uint64 `json:"build_number"`
	Major       uint64 `json:"major"`
	Minor       uint64 `json:"minor"`
}
