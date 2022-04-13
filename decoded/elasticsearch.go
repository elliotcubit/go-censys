package decoded

// Hopefully, one doesn't need to use these struct names... :P

type Elasticsearch struct {
	HttpInfo   *HttpInfo                 `json:"http_info"`
	SystemInfo *Elasticsearch_SystemInfo `json:"system_info"`
	NodeInfo   *Elasticsearch_NodeInfo   `json:"node_info"`
}

type Elasticsearch_SystemInfo struct {
	Name        *string                           `json:"name"`
	ClusterUUID *string                           `json:"cluster_uuid"`
	Version     *Elasticsearch_SystemInfo_Version `json:"version"`
	Tagline     *string                           `json:"tagline"`
}

type Elasticsearch_SystemInfo_Version struct {
	Number           *string `json:"number"`
	BuildFlavor      *string `json:"build_flavor"`
	BuildType        *string `json:"build_type"`
	BuildHash        *string `json:"build_hash"`
	BuildDate        *string `json:"build_date"`
	BuildSnapshot    *bool   `json:"build_snapshot"`
	LuceneVersion    *string `json:"lucene_version"`
	MinWireCompatVer *string `json:"min_wire_compat_ver"`
	MaxIDXCompatVer  *string `json:"max_idx_compat_ver"`
}

type Elasticsearch_NodeInfo struct {
	ClusterCombinedInfo *Elasticsearch_NodeInfo_ClusterCombinedInfo `json:"cluster_combined_info"`
	Nodes               []Elasticsearch_NodeInfo_Node               `json:"nodes"`
}

type Elasticsearch_NodeInfo_ClusterCombinedInfo struct {
	Name       *string                                                `json:"name"`
	UUID       *string                                                `json:"uuid"`
	Timestamp  *uint64                                                `json:"timestamp"`
	Status     *string                                                `json:"status"`
	Filesystem *Elasticsearch_NodeInfo_ClusterCombinedInfo_Filesystem `json:"filesystem"`
}

type Elasticsearch_NodeInfo_ClusterCombinedInfo_Filesystem struct {
	Available        *string `json:"available"`
	AvailableInBytes *uint64 `json:"available_in_bytes"`
	Free             *string `json:"free"`
	FreeInBytes      *uint64 `json:"free_in_bytes"`
	Total            *string `json:"total"`
	TotalInBytes     *string `json:"total_in_bytes"`
}

type Elasticsearch_NodeInfo_ClusterCombinedInfo_Indices struct {
	Count *uint64                                                   `json:"count"`
	Docs  *Elasticsearch_NodeInfo_ClusterCombinedInfo_Indices_Docs  `json:"docs"`
	Store *Elasticsearch_NodeInfo_ClusterCombinedInfo_Indices_Store `json:"store"`
}

type Elasticsearch_NodeInfo_ClusterCombinedInfo_Indices_Docs struct {
	Count   *uint64 `json:"count"`
	Deleted *uint64 `json:"deleted"`
}

type Elasticsearch_NodeInfo_ClusterCombinedInfo_Indices_Store struct {
	SizeInBytes     *uint64 `json:"size_in_bytes"`
	ReservedInBytes *uint64 `json:"reserved_in_bytes"`
}

type Elasticsearch_NodeInfo_Node struct {
	Name     *string                               `json:"name"`
	NodeData *Elasticsearch_NodeInfo_Node_NodeData `json:"node_data"`
}

type Elasticsearch_NodeInfo_Node_NodeData struct {
	Name                *string                                           `json:"name"`
	Host                *string                                           `json:"host"`
	IP                  *string                                           `json:"ip"`
	Version             *string                                           `json:"version"`
	BuildFlavor         *string                                           `json:"build_flavor"`
	BuildType           *string                                           `json:"build_type"`
	BuildHash           *string                                           `json:"build_hash"`
	TotalIndexingBuffer *uint64                                           `json:"total_indexing_buffer"`
	Roles               []string                                          `json:"roles"`
	Settings            *Elasticsearch_NodeInfo_Node_NodeData_Settings    `json:"settings"`
	OS                  *Elasticsearch_NodeInfo_Node_NodeData_OS          `json:"os"`
	JVM                 *Elasticsearch_NodeInfo_Node_NodeData_JVM         `json:"jvm"`
	ThreadPoolList      []Elasticsearch_NodeInfo_Node_NodeData_ThreadPool `json:"thread_pool_list"`
	Modules             []Elasticsearch_NodeInfo_Node_NodeData_Module     `json:"modules"`
	IngestProcessors    []string                                          `json:"ingest_processors"`
}

type Elasticsearch_NodeInfo_Node_NodeData_JVM struct {
	Version         *string  `json:"version"`
	VMName          *string  `json:"vm_name"`
	VMVersion       *string  `json:"vm_version"`
	VMVendor        *string  `json:"vm_vendor"`
	StartTime       *string  `json:"start_time"`
	StartTimeMillis *uint64  `json:"start_time_ms"`
	GC              []string `json:"gc"`
	MemoryPools     []string `json:"memory_pools"`
	InputArgs       []string `json:"input_args"`
}

type Elasticsearch_NodeInfo_Node_NodeData_Module struct {
	Name           *string  `json:"name"`
	Version        *string  `json:"version"`
	ElasticVersion *string  `json:"elastic_version"`
	JavaVersion    *string  `json:"java_version"`
	Desc           *string  `json:"desc"`
	ClassName      *string  `json:"class_name"`
	ExtPlugins     []string `json:"ext_plugins"`
	HasNativeCtrl  *bool    `json:"has_native_ctrl"`
}

type Elasticsearch_NodeInfo_Node_NodeData_OS struct {
	RefreshIntervalMillis *uint64 `json:"refresh_interval_ms"`
	Name                  *string `json:"name"`
	PrettyName            *string `json:"pretty_name"`
	Arch                  *string `json:"arch"`
	Version               *string `json:"version"`
	AvailableProc         *int    `json:"available_proc"`
	AllocatedProc         *int    `json:"allocated_proc"`
}

type Elasticsearch_NodeInfo_Node_NodeData_Settings struct {
	ClusterName *string                                             `json:"cluster_name"`
	Node        *Elasticsearch_NodeInfo_Node_NodeData_Settings_Node `json:"node"`
}

type Elasticsearch_NodeInfo_Node_NodeData_ThreadPool struct {
	Type      *string `json:"type"`
	Min       *int    `json:"min"`
	Max       *int    `json:"max"`
	KeepAlive *string `json:"keep_alive"`
	QueueSize *int    `json:"queue_size"`
}

type Elasticsearch_NodeInfo_Node_NodeData_Settings_Node struct {
	Attr *Elasticsearch_NodeInfo_Node_NodeData_Settings_Node_Attr `json:"attr"`
	Name *string                                                  `json:"name"`
}

type Elasticsearch_NodeInfo_Node_NodeData_Settings_Node_Attr struct {
	XPackInstalled *string                                                     `json:"xpack_installed"`
	ML             *Elasticsearch_NodeInfo_Node_NodeData_Settings_Node_Attr_ML `json:"ml"`
}

type Elasticsearch_NodeInfo_Node_NodeData_Settings_Node_Attr_ML struct {
	MachineMemory *string `json:"machine_memory"`
	MaxOpenJobs   *string `json:"max_open_jobs"`
	Enabled       *string `json:"enabled"`
}
