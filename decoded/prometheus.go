package decoded

type Prometheus struct {
	HttpInfo Prometheus_HttpInfo `json:"http_info"`
	Response Prometheus_Response `json:"response"`
}

type Prometheus_Response struct {
	ActiveTargets      []Prometheus_Response_ActiveTargets    `json:"active_targets"`
	AllVersions        []string                               `json:"all_versions"`
	DroppedTargets     []Prometheus_Response_DroppedTargets   `json:"dropped_targets"`
	GoVersions         string                                 `json:"go_versions"`
	PrometheusVersions Prometheus_Response_PrometheusVersions `json:"prometheus_versions"`
}

type Prometheus_Response_PrometheusVersions struct {
	GoVersion string `json:"go_version"`
	Revision  string `json:"revision"`
	Version   string `json:"version"`
}

type Prometheus_Response_DroppedTargets struct {
	Address     string `json:"address"`
	Job         string `json:"job"`
	MetricsPath string `json:"metrics_path"`
	Scheme      string `json:"scheme"`
}

type Prometheus_Response_ActiveTargets struct {
	DiscoveredLabels Prometheus_Response_ActiveTargets_DiscoveredLabels `json:"discovered_labels"`
	Health           string                                             `json:"health"`
	Labels           Prometheus_Response_ActiveTargets_Labels           `json:"labels"`
	LastError        string                                             `json:"last_error"`
	LastScrape       string                                             `json:"last_scrape"`
	ScrapeUrl        string                                             `json:"scrape_url"`
}

type Prometheus_Response_ActiveTargets_Labels struct {
	Instance string `json:"instance"`
	Job      string `json:"job"`
}

type Prometheus_Response_ActiveTargets_DiscoveredLabels struct {
	Address     string `json:"address"`
	Job         string `json:"job"`
	MetricsPath string `json:"metrics_path"`
	Scheme      string `json:"scheme"`
}

type Prometheus_HttpInfo struct {
	Headers    HttpHeaders `json:"headers"`
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
}
