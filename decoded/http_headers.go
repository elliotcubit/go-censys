package decoded

type HttpHeaders struct {
	AcceptPatch              []string `json:"accept_patch"`
	AcceptRanges             []string `json:"accept_ranges"`
	AccessControlAllowOrigin []string `json:"access_control_allow_origin"`
	Age                      []string `json:"age"`
	Allow                    []string `json:"allow"`
	AltSvc                   []string `json:"alt_svc"`
	AlternateProtocol        []string `json:"alternate_protocol"`
	CacheControl             []string `json:"cache_control"`
	Connection               []string `json:"connection"`
	ContentDisposition       []string `json:"content_disposition"`
	ContentEncoding          []string `json:"content_encoding"`
	ContentLanguage          []string `json:"content_language"`
	ContentLength            []string `json:"content_length"`
	ContentLocation          []string `json:"content_location"`
	ContentMD5               []string `json:"content_md5"`
	ContentRange             []string `json:"content_range"`
	ContentSecurityPolicy    []string `json:"content_security_policy"`
	ContentType              []string `json:"content_type"`
	Date                     []string `json:"date"`
	ETag                     []string `json:"etag"`
	Expires                  []string `json:"expires"`
	LastModified             []string `json:"last_modified"`
	Link                     []string `json:"link"`
	Location                 []string `json:"location"`
	P3P                      []string `json:"p3p"`
	Pragma                   []string `json:"pragma"`
	ProxyAgent               []string `json:"proxy_agent"`
	ProxyAuthenticate        []string `json:"proxy_authenticate"`
	PublicKeyPins            []string `json:"public_key_pins"`
	// Yes - the HTTP header "referer" is a misspelling of "referrer."
	// See RFC7231 5.5.2
	Referer                 []string            `json:"referer"`
	Refresh                 []string            `json:"refresh"`
	RetryAfter              []string            `json:"retry_after"`
	Server                  []string            `json:"server"`
	SetCookie               []string            `json:"set_cookie"`
	Status                  []string            `json:"status"`
	StrictTransportSecurity []string            `json:"strict_transport_security"`
	Trailer                 []string            `json:"trailer"`
	TransferEncoding        []string            `json:"transfer_encoding"`
	Unknown                 map[string][]string `json:"unknown"`
	Upgrade                 []string            `json:"upgrade"`
	Vary                    []string            `json:"vary"`
	Via                     []string            `json:"via"`
	Warning                 []string            `json:"warning"`
	WWWAuthenticate         []string            `json:"www_authenticate"`
	XContentDuration        []string            `json:"x_content_duration"`
	XContentSecurityPolicy  []string            `json:"x_content_security_policy"`
	XContentTypeOptions     []string            `json:"x_content_type_options"`
	XForwardedFor           []string            `json:"x_forwarded_for"`
	XFrameOptions           []string            `json:"x_frame_options"`
	XPoweredBy              []string            `json:"x_powered_by"`
	XRealIP                 []string            `json:"x_real_ip"`
	XUACompatible           []string            `json:"x_ua_compatible"`
	XWebkitCSP              []string            `json:"x_webkit_csp"`
	XXSSProtection          []string            `json:"x_xss_protection"`
}
