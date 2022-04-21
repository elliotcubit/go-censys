package decoded

type Mysql struct {
	AuthPluginData  string                `json:"auth_plugin_data"`
	AuthPluginName  string                `json:"auth_plugin_name"`
	CapabilityFlags Mysql_CapabilityFlags `json:"capability_flags"`
	CharacterSet    uint64                `json:"character_set"`
	ConnectionId    uint64                `json:"connection_id"`
	ErrorCode       int64                 `json:"error_code"`
	ErrorId         string                `json:"error_id"`
	ErrorMessage    string                `json:"error_message"`
	ProtocolVersion uint64                `json:"protocol_version"`
	ServerVersion   string                `json:"server_version"`
	StatusFlags     Mysql_StatusFlags     `json:"status_flags"`
}

type Mysql_StatusFlags struct {
	ServerMoreResultsExists        bool `json:"server_more_results_exists"`
	ServerPsOutParams              bool `json:"server_ps_out_params"`
	ServerQueryNoGoodIndexUsed     bool `json:"server_query_no_good_index_used"`
	ServerQueryNoIndexUsed         bool `json:"server_query_no_index_used"`
	ServerQueryWasSlow             bool `json:"server_query_was_slow"`
	ServerSessionStateChanged      bool `json:"server_session_state_changed"`
	ServerStatusAutocommit         bool `json:"server_status_autocommit"`
	ServerStatusCursorExists       bool `json:"server_status_cursor_exists"`
	ServerStatusDbDropped          bool `json:"server_status_db_dropped"`
	ServerStatusInTrans            bool `json:"server_status_in_trans"`
	ServerStatusInTransReadonly    bool `json:"server_status_in_trans_readonly"`
	ServerStatusLastRowSent        bool `json:"server_status_last_row_sent"`
	ServerStatusMetadataChanged    bool `json:"server_status_metadata_changed"`
	ServerStatusNoBackslashEscapes bool `json:"server_status_no_backslash_escapes"`
}

type Mysql_CapabilityFlags struct {
	ClientCanHandleExpiredPasswords  bool `json:"client_can_handle_expired_passwords"`
	ClientCompress                   bool `json:"client_compress"`
	ClientConnectAttrs               bool `json:"client_connect_attrs"`
	ClientConnectWithDb              bool `json:"client_connect_with_db"`
	ClientDeprecatedEof              bool `json:"client_deprecated_eof"`
	ClientFoundRows                  bool `json:"client_found_rows"`
	ClientIgnoreSigpipe              bool `json:"client_ignore_sigpipe"`
	ClientIgnoreSpace                bool `json:"client_ignore_space"`
	ClientInteractive                bool `json:"client_interactive"`
	ClientLocalFiles                 bool `json:"client_local_files"`
	ClientLongFlag                   bool `json:"client_long_flag"`
	ClientLongPassword               bool `json:"client_long_password"`
	ClientMultiResults               bool `json:"client_multi_results"`
	ClientMultiStatements            bool `json:"client_multi_statements"`
	ClientNoSchema                   bool `json:"client_no_schema"`
	ClientOdbc                       bool `json:"client_odbc"`
	ClientPluginAuth                 bool `json:"client_plugin_auth"`
	ClientPluginAuthLenEncClientData bool `json:"client_plugin_auth_len_enc_client_data"`
	ClientProtocol41                 bool `json:"client_protocol_41"`
	ClientPsMultiResults             bool `json:"client_ps_multi_results"`
	ClientReserved                   bool `json:"client_reserved"`
	ClientSecureConnection           bool `json:"client_secure_connection"`
	ClientSessionTrack               bool `json:"client_session_track"`
	ClientSsl                        bool `json:"client_ssl"`
	ClientTransactions               bool `json:"client_transactions"`
}
