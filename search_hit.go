package censys

type SearchHostsResult struct {
	// The Query that produced this result
	Query string `json:"query"`

	// The Total number of hits
	Total int `json:"total"`

	// The actual results
	Hits []SearchHit `json:"hits"`

	// Links to the prev/next pages of results
	Links SearchResultLinks `json:"links"`
}

type SearchResultLinks struct {
	// Prev is a token that can be passed to get the previous page of search results.
	Prev string `json:"prev"`

	// Next us a token that can be passed to get the next page of search results.
	Next string `json:"next"`
}

type SearchHit struct {
	// IP is the IP address of the host.
	IP string `json:"ip"`

	// Name is the Name of the host if it is a virtual host, which is provided
	// in the TLS handshake in the Server Name Indication extension
	// (as well as any other place that a domain name could be used at scan time)
	Name string `json:"name"`

	// Services is a list of all Services that were observed on the host.
	// Note that this list contains only partial information about the services.
	Services []SmallService `json:"services"`

	// Location is location of the host.
	Location Location `json:"location"`

	// AutonomousSystem countains routing information for the host.
	AutonomousSystem AutonomousSystem `json:"autonomous_system"`
}

// SmallService is a Service with most fields removed.
// It's the shortened verison, returned in search results.
type SmallService struct {
	// Port is the port number the service was discovered on.
	Port uint16 `json:"port"`

	// ServiceName is the name/type of the service. Note that this slightly differs
	// from the meaning of `DecodedType`. `DecodedType` is merely the format of the
	// decoded data field, whereas `ServiceName` definitively describes the type
	// of service. This is most relevant for some services, such as `MURMUR`, which
	// use the same `DecodedType` as `UNKNOWN` services.
	ServiceName string `json:"service_name"`

	// TransportProtocol is the OSI Layer 4 protocol over which this Service's scan
	// was conducted. See `TransportProtocolTCP`/etc for potential values.
	TransportProtocol string `json:"transport_protocol"`

	// Certificate is the SHA256 hash of the Service's TLS
	// Certificate, if it is presenting one
	Certificate string `json:"certificate"`
}
