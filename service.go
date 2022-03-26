package censys

import (
	"encoding/json"
	"time"

	"github.com/elliotcubit/go-censys/decoded"
)

const (
	// DecodedTypeAmqp indicates a Service is Amqp
	DecodedTypeAmqp = "amqp"

	// DecodedTypeAnyConnect indicates a Service is AnyConnect
	DecodedTypeAnyConnect = "any_connect"

	// DecodedTypeBacnet indicates a Service is Bacnet
	DecodedTypeBacnet = "bacnet"

	// DecodedTypeCoap indicates a Service is Coap
	DecodedTypeCoap = "coap"

	// DecodedTypeCwmp indicates a Service is Cwmp
	DecodedTypeCwmp = "cwmp"

	// DecodedTypeDnp3 indicates a Service is Dnp3
	DecodedTypeDnp3 = "dnp3"

	// DecodedTypeDns indicates a Service is Dns.
	DecodedTypeDns = "dns"
)

const (
	PerspectiveHE      = "PERSPECTIVE_HE"
	PerspectiveNTT     = "PERSPECTIVE_NTT"
	PerspectiveTATA    = "PERSPECTIVE_TATA"
	PerspectiveORANGE  = "PERSPECTIVE_ORANGE"
	PerspectiveTELIA   = "PERSPECTIVE_TELIA"
	PerspectiveUnknown = "PERSPECTIVE_UNKNOWN"
)

const (
	// TransportProtocolTCP indicates the scan took place over TCP
	TransportProtocolTCP = "TCP"

	// TransportProtocolUDP indicates the scan took place over UDP
	TransportProtocolUDP = "UDP"

	// TransportProtocolQUIC indicates the scan took place over QUIC
	// Note that QUIC scans also, technically, occur over UDP since QUIC is
	// implemented on top of the L4 UDP protocol.
	TransportProtocolQUIC = "QUIC"
)

// Service represents a service living on a host. Services run on a port,
// over a transport protocol, and perform a certain function. Censys extracts
// different types of information from every unique service type it supports.
type Service struct {
	// Banner is (generally) the initial response of the Service. It can be seen
	// as a "preview" of the service. For services that are deeply parsed
	// (i.e. not `UNKNOWN`), this field is only populated when the protocol is
	// banner-based or has a meaningful "banner" immediately available on connection;
	// e.g. FTP, HTTP, SMTP.
	Banner string `json:"banner"`

	// Certificate is the SHA256 hash of the Service's TLS
	// Certificate, if it is presenting one.
	Certificate string `json:"certificate"`

	// DecodedType can be used to identify which Decoded field, below, will be
	// non-nil. Use the constants `DecodedTypeDNS`/etc for potential values.
	DecodedType string `json:"_decoded"`

	// ExtendedServiceName is the "full" / "canonical" name of the service.
	// For example, `HTTP` services are always the decoded type `DecodedTypeHTTP`,
	// and the `ServiceName` is always `HTTP`, but generally `HTTP` served over
	// TLS is called `HTTPS`; that is the value that will be present here.
	//
	// Some services, such as `SMTP`, will also have the pattern `SMTP-STARTTLS`,
	// indicating that TLS was negotiated via the STARTTLS command rather than
	// through implicit TLS.
	//
	// Note that generally, this field only differs from the `ServiceName` when
	// there is a another common-use name for the protocol. `UPNP` served over TLS,
	// for example, does not have an `ExtendedServiceName` of `UPNPS`.
	ExtendedServiceName string `json:"extended_service_name"`

	// ObservedAt is the time this Service's scan occurred.
	ObservedAt time.Time `json:"observed_at"`

	// PerspectiveID is the ID of the scan perspective from which this service's
	// scan ocurred. Use the constants `PerspectiveHE`... for potential values.
	PerspectiveID string `json:"perspective_id"`

	// Port is the port number the service was discovered on.
	Port uint16 `json:"port"`

	// ServiceName is the name/type of the service. Note that this slightly differs
	// from the meaning of `DecodedType`. `DecodedType` is merely the format of the
	// decoded data field, whereas `ServiceName` definitively describes the type
	// of service. This is most relevant for some services, such as `MURMUR`, which
	// use the same `DecodedType` as `UNKNOWN` services.
	ServiceName string `json:"service_name"`

	// Software is a list of Software suspected to be running on the Service.
	// For example, a Service might have a Software entry for nginx.
	Software []Software `json:"software"`

	// SourceIP is the IP address from which this Service's scan originated.
	SourceIP string `json:"source_ip"`

	// TLS contains the TLS certificate of the service, if it is presenting one.
	TLS *json.RawMessage `json:"tls"`

	// TransportProtocol is the OSI Layer 4 protocol over which this Service's scan
	// was conducted. See `TransprotProtocolTCP`/etc for potential values.
	TransportProtocol string `json:"transport_protocol"`

	// Truncated, if set, indicates that the service has had its deep scan data
	// truncated. This occurs when a large amount of services are on the same host.
	//
	// If a Service is `Truncated`, it will not have a `Decoded`!
	Truncated bool `json:"truncated"`

	// Amqp contains information about an AMQP service.
	// It is non-nil when DecodedType is DecodedTypeAmqp.
	Amqp *decoded.Amqp `json:"amqp"`

	// AnyConnect contains information about an AnyConnect service.
	// It is non-nil when DecodedType is DecodedTypeAnyConnect.
	AnyConnect *decoded.AnyConnect `json:"any_connect"`

	// Bacnet contains information about an Bacnet service.
	// It is non-nil when DecodedType is DecodedTypeBacnet.
	Bacnet *decoded.Bacnet `json:"bacnet"`

	// Coap contains information about an Coap service.
	// It is non-nil when DecodedType is DecodedTypeCoap.
	Coap *decoded.Coap `json:"coap"`

	// Cwmp contains information about an Cwmp service.
	// It is non-nil when DecodedType is DecodedTypeCwmp.
	Cwmp *decoded.Cwmp `json:"cwmp"`

	// Dnp3 contains information about an Dnp3 service.
	// It is non-nil when DecodedType is DecodedTypeDnp3.
	Dnp3 *decoded.Dnp3 `json:"dnp3"`

	// Dns contains information about a DNS service.
	// It is non-nil when DecodedType is DecodedTypeDns.
	Dns *decoded.Dns `json:"dns"`
}
