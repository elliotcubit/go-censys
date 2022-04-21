package decoded

type Postgres struct {
	AuthenticationMode Postgres_AuthenticationMode `json:"authentication_mode"`
	ProtocolError      Postgres_ProtocolError      `json:"protocol_error"`
	StartupError       Postgres_StartupError       `json:"startup_error"`
	SupportedVersions  []string                    `json:"supported_versions"`
	TransactionStatus  string                      `json:"transaction_status"`
}

type Postgres_StartupError struct {
	UnknownErrorTag  string `json:"_unknown_error_tag"`
	Code             string `json:"code"`
	Constraint       string `json:"constraint"`
	Data             string `json:"data"`
	Detail           string `json:"detail"`
	File             string `json:"file"`
	Hint             string `json:"hint"`
	InternalPosition string `json:"internal_position"`
	InternalQuery    string `json:"internal_query"`
	Line             string `json:"line"`
	Message          string `json:"message"`
	Position         string `json:"position"`
	Routine          string `json:"routine"`
	Schema           string `json:"schema"`
	Severity         string `json:"severity"`
	SeverityV        string `json:"severity_v"`
	Table            string `json:"table"`
	Where            string `json:"where"`
}

type Postgres_ProtocolError struct {
	UnknownErrorTag  string `json:"_unknown_error_tag"`
	Code             string `json:"code"`
	Constraint       string `json:"constraint"`
	Data             string `json:"data"`
	Detail           string `json:"detail"`
	File             string `json:"file"`
	Hint             string `json:"hint"`
	InternalPosition string `json:"internal_position"`
	InternalQuery    string `json:"internal_query"`
	Line             string `json:"line"`
	Message          string `json:"message"`
	Position         string `json:"position"`
	Routine          string `json:"routine"`
	Schema           string `json:"schema"`
	Severity         string `json:"severity"`
	SeverityV        string `json:"severity_v"`
	Table            string `json:"table"`
	Where            string `json:"where"`
}

type Postgres_AuthenticationMode struct {
	Mode    string `json:"mode"`
	Payload string `json:"payload"`
}
