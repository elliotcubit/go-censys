package decoded

type Mongodb struct {
	BuildInfo Mongodb_BuildInfo `json:"build_info"`
	IsMaster  Mongodb_IsMaster  `json:"is_master"`
}

type Mongodb_IsMaster struct {
	IsMaster                     bool `json:"is_master"`
	LogicalSessionTimeoutMinutes int  `json:"logical_session_timeout_minutes"`
	MaxBsonObjectSize            int  `json:"max_bson_object_size"`
	MaxMessageSizeBytes          int  `json:"max_message_size_bytes"`
	MaxWireVersion               int  `json:"max_wire_version"`
	MaxWriteBatchSize            int  `json:"max_write_batch_size"`
	MinWireVersion               int  `json:"min_wire_version"`
	ReadOnly                     bool `json:"read_only"`
}

type Mongodb_BuildInfo struct {
	BuildEnvironment Mongodb_BuildInfo_BuildEnvironment `json:"build_environment"`
	GitVersion       string                             `json:"git_version"`
	Version          string                             `json:"version"`
}

type Mongodb_BuildInfo_BuildEnvironment struct {
	Cc         string `json:"cc"`
	CcFlags    string `json:"cc_flags"`
	Cxx        string `json:"cxx"`
	CxxFlags   string `json:"cxx_flags"`
	DistArch   string `json:"dist_arch"`
	DistMod    string `json:"dist_mod"`
	LinkFlags  string `json:"link_flags"`
	TargetArch string `json:"target_arch"`
	TargetOs   string `json:"target_os"`
}
