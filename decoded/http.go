package decoded

type Http struct {
	Request       Http_Request  `json:"request"`
	Response      Http_Response `json:"response"`
	SupportsHttp2 bool          `json:"supports_http2"`
}

type Http_Response struct {
	Body         string                  `json:"body"`
	BodyHash     string                  `json:"body_hash"`
	BodySize     int                     `json:"body_size"`
	Favicons     []Http_Response_Favicon `json:"favicons"`
	Headers      HttpHeaders             `json:"headers"`
	HtmlTags     []string                `json:"html_tags"`
	HtmlTitle    string                  `json:"html_title"`
	Protocol     string                  `json:"protocol"`
	StatusCode   int                     `json:"status_code"`
	StatusReason string                  `json:"status_reason"`
}

type Http_Response_Favicon struct {
	Md5Hash string `json:"md5_hash"`
	Name    string `json:"name"`
	Size    int    `json:"size"`
}

type Http_Request struct {
	Body    string      `json:"body"`
	Headers HttpHeaders `json:"headers"`
	Method  string      `json:"method"`
	Uri     string      `json:"uri"`
}
