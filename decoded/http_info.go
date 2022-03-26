package decoded

type HttpInfo struct {
	Body         string      `json:"body"`
	BodyHash     string      `json:"body_hash"`
	BodySize     int         `json:"body_size"`
	Favicons     []Favicon   `json:"favicons"`
	Headers      HttpHeaders `json:"headers"`
	HTMLTags     []string    `json:"html_tags"`
	HTMLTitle    string      `json:"html_title"`
	Protocol     string      `json:"protocol"`
	StatusCode   int         `json:"status_code"`
	StatusReason string      `json:"status_reason"`
}
