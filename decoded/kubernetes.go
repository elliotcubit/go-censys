package decoded

type Kubernetes struct {
	Endpoints                []Kubernetes_Endpoint  `json:"endpoints"`
	KubernetesDashboardFound bool                   `json:"kubernetes_dashboard_found"`
	Nodes                    []Kubernetes_Node      `json:"nodes"`
	PodNames                 string                 `json:"pod_names"`
	Roles                    []Kubernetes_Role      `json:"roles"`
	VersionInfo              Kubernetes_VersionInfo `json:"version_info"`
}

type Kubernetes_VersionInfo struct {
	BuildDate    string `json:"build_date"`
	Compiler     string `json:"compiler"`
	GitCommit    string `json:"git_commit"`
	GitTreeState string `json:"git_tree_state"`
	GitVersion   string `json:"git_version"`
	GoVersion    string `json:"go_version"`
	Major        string `json:"major"`
	Minor        string `json:"minor"`
	Platform     string `json:"platform"`
}

type Kubernetes_Role struct {
	Name  string                 `json:"name"`
	Rules []Kubernetes_Role_Rule `json:"rules"`
}

type Kubernetes_Role_Rule struct {
	ApiGroups []string `json:"api_groups"`
	Resources []string `json:"resources"`
	Verbs     []string `json:"verbs"`
}

type Kubernetes_Node struct {
	Addresses               []Kubernetes_Node_Address `json:"addresses"`
	Architecture            string                    `json:"architecture"`
	ContainerRuntimeVersion string                    `json:"container_runtime_version"`
	Images                  string                    `json:"images"`
	KernelVersion           string                    `json:"kernel_version"`
	KubeProxyVersion        string                    `json:"kube_proxy_version"`
	KubeletVersion          string                    `json:"kubelet_version"`
	Name                    string                    `json:"name"`
	OperatingSystem         string                    `json:"operating_system"`
	OsImage                 string                    `json:"os_image"`
}

type Kubernetes_Node_Address struct {
	Address     string `json:"address"`
	AddressType string `json:"address_type"`
}

type Kubernetes_Endpoint struct {
	Name     string                       `json:"name"`
	SelfLink string                       `json:"self_link"`
	Subsets  []Kubernetes_Endpoint_Subset `json:"subsets"`
}

type Kubernetes_Endpoint_Subset struct {
	Addresses []Kubernetes_Endpoint_Subset_Address `json:"addresses"`
	Ports     []Kubernetes_Endpoint_Subset_Port    `json:"ports"`
}

type Kubernetes_Endpoint_Subset_Port struct {
	Name     string `json:"name"`
	Port     uint64 `json:"port"`
	Protocol string `json:"protocol"`
}

type Kubernetes_Endpoint_Subset_Address struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
	NodeName string `json:"node_name"`
}
