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

	// DecodedTypeElasticsearch indicates a Service is Elasticsearch.
	DecodedTypeElasticsearch = "elasticsearch"

	// DecodedTypeFox indicates a Service is Fox.
	DecodedTypeFox = "fox"

	// DecodedTypeFtp indicates a Service is Ftp.
	DecodedTypeFtp = "ftp"

	// DecodedTypeIke indicates a Service is Ike.
	DecodedTypeIke = "ike"

	// DecodedTypeImap indicates a Service is Imap.
	DecodedTypeImap = "imap"

	// DecodedTypeIpmi indicates a Service is Ipmi.
	DecodedTypeIpmi = "ipmi"

	// DecodedTypeIpp indicates a Service is Ipp.
	DecodedTypeIpp = "ipp"

	// DecodedTypeKubernetes indicates a Service is Kubernetes.
	DecodedTypeKubernetes = "kubernetes"

	// DecodedTypeLdap indicates a Service is Ldap.
	DecodedTypeLdap = "ldap"

	// DecodedTypeMemcached indicates a Service is Memcached.
	DecodedTypeMemcached = "memcached"

	// DecodedTypeMms indicates a Service is Mms.
	DecodedTypeMms = "mms"

	// DecodedTypeModbus indicates a Service is Modbus.
	DecodedTypeModbus = "modbus"

	// DecodedTypeMongodb indicates a Service is Mongodb.
	DecodedTypeMongodb = "mongodb"

	// DecodedTypeMqtt indicates a Service is Mqtt.
	DecodedTypeMqtt = "mqtt"

	// DecodedTypeMssql indicates a Service is Mssql.
	DecodedTypeMssql = "mssql"

	// DecodedTypeMysql indicates a Service is Mysql.
	DecodedTypeMysql = "mysql"

	// DecodedTypeNtp indicates a Service is Ntp.
	DecodedTypeNtp = "ntp"

	// DecodedTypeOpenvpn indicates a Service is Openvpn.
	DecodedTypeOpenvpn = "openvpn"

	// DecodedTypeOracle indicates a Service is Oracle.
	DecodedTypeOracle = "oracle"

	// DecodedTypePcAnywhere indicates a Service is PcAnywhere.
	DecodedTypePcAnywhere = "pc_anywhere"

	// DecodedTypePop3 indicates a Service is Pop3.
	DecodedTypePop3 = "pop3"

	// DecodedTypePostgres indicates a Service is Postgres.
	DecodedTypePostgres = "postgres"

	// DecodedTypePptp indicates a Service is Pptp.
	DecodedTypePptp = "pptp"

	// DecodedTypePrometheus indicates a Service is Prometheus.
	DecodedTypePrometheus = "prometheus"

	// DecodedTypeRdp indicates a Service is Rdp.
	DecodedTypeRdp = "rdp"

	// DecodedTypeRedis indicates a Service is Redis.
	DecodedTypeRedis = "redis"

	// DecodedTypeS7 indicates a Service is S7.
	DecodedTypeS7 = "s7"

	// DecodedTypeSip indicates a Service is Sip.
	DecodedTypeSip = "sip"

	// DecodedTypeSkinny indicates a Service is Skinny.
	DecodedTypeSkinny = "skinny"

	// DecodedTypeSmtp indicates a Service is Smtp.
	DecodedTypeSmtp = "smtp"

	// DecodedTypeSnmp indicates a Service is Snmp.
	DecodedTypeSnmp = "snmp"

	// DecodedTypeSsdp indicates a Service is Ssdp.
	DecodedTypeSsdp = "ssdp"

	// DecodedTypeSsh indicates a Service is Ssh.
	DecodedTypeSsh = "ssh"

	// DecodedTypeTeamViewer indicates a Service is TeamViewer.
	DecodedTypeTeamViewer = "team_viewer"

	// DecodedTypeTelnet indicates a Service is Telnet.
	DecodedTypeTelnet = "telnet"

	// DecodedTypeUpnp indicates a Service is Upnp.
	DecodedTypeUpnp = "upnp"

	// DecodedTypeVnc indicates a Service is Vnc.
	DecodedTypeVnc = "vnc"

	// DecodedTypeX11 indicates a Service is X11.
	DecodedTypeX11 = "x11"
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

	// Bacnet contains information about a Bacnet service.
	// It is non-nil when DecodedType is DecodedTypeBacnet.
	Bacnet *decoded.Bacnet `json:"bacnet"`

	// Coap contains information about a Coap service.
	// It is non-nil when DecodedType is DecodedTypeCoap.
	Coap *decoded.Coap `json:"coap"`

	// Cwmp contains information about a Cwmp service.
	// It is non-nil when DecodedType is DecodedTypeCwmp.
	Cwmp *decoded.Cwmp `json:"cwmp"`

	// Dnp3 contains information about a Dnp3 service.
	// It is non-nil when DecodedType is DecodedTypeDnp3.
	Dnp3 *decoded.Dnp3 `json:"dnp3"`

	// Dns contains information about a DNS service.
	// It is non-nil when DecodedType is DecodedTypeDns.
	Dns *decoded.Dns `json:"dns"`

	// Elasticsearch contains information about an Elasticsearch service.
	// It is non-nil when DecodedType is DecodedTypeElasticsearch.
	Elasticsearch *decoded.Elasticsearch `json:"elasticsearch"`

	// Fox contains information about an Fox service.
	// It is non-nil when DecodedType is DecodedTypeFox.
	Fox *decoded.Elasticsearch `json:"fox"`

	// Ftp contains information about an Ftp service.
	// It is non-nil when DecodedType is DecodedTypeFtp.
	Ftp *decoded.Ftp `json:"ftp"`

	// Http contains information about an Http service.
	// It is non-nil when DecodedType is DecodedTypeHttp.
	Http *decoded.Http `json:"http"`

	// Ike contains information about an Ike service.
	// It is non-nil when DecodedType is DecodedTypeIke.
	Ike *decoded.Ike `json:"ike"`

	// Imap contains information about an Imap service.
	// It is non-nil when DecodedType is DecodedTypeImap.
	Imap *decoded.Imap `json:"imap"`

	// Ipmi contains information about an Ipmi service.
	// It is non-nil when DecodedType is DecodedTypeIpmi.
	Ipmi *decoded.Ipmi `json:"ipmi"`

	// Ipp contains information about an Ipp service.
	// It is non-nil when DecodedType is DecodedTypeIpp.
	Ipp *decoded.Ipp `json:"ipp"`

	// Kubernetes contains information about an Kubernetes service.
	// It is non-nil when DecodedType is DecodedTypeKubernetese
	Kubernetes *decoded.Kubernetes `json:"kubernetes"`

	// Ldap contains information about an Ldap service.
	// It is non-nil when DecodedType is DecodedTypeLdap
	Ldap *decoded.Ldap `json:"ldap"`

	// Memcached contains information about an Memcached service.
	// It is non-nil when DecodedType is DecodedTypeMemcached
	Memcached *decoded.Memcached `json:"memcached"`

	// Mms contains information about an Mms service.
	// It is non-nil when DecodedType is DecodedTypeMms
	Mms *decoded.Mms `json:"mms"`

	// Modbus contains information about an Modbus service.
	// It is non-nil when DecodedType is DecodedTypeModbus
	Modbus *decoded.Modbus `json:"modbus"`

	// Mongodb contains information about an Mongodb service.
	// It is non-nil when DecodedType is DecodedTypeMongodb
	Mongodb *decoded.Mongodb `json:"mongodb"`

	// Mqtt contains information about an Mqtt service.
	// It is non-nil when DecodedType is DecodedTypeMqtt
	Mqtt *decoded.Mqtt `json:"mqtt"`

	// Mssql contains information about an Mssql service.
	// It is non-nil when DecodedType is DecodedTypeMssql
	Mssql *decoded.Mssql `json:"mssql"`

	// Mysql contains information about an Mysql service.
	// It is non-nil when DecodedType is DecodedTypeMysql
	Mysql *decoded.Mysql `json:"mysql"`

	// Ntp contains information about an Ntp service.
	// It is non-nil when DecodedType is DecodedTypeNtp
	Ntp *decoded.Ntp `json:"ntp"`

	// Openvpn contains information about an Openvpn service.
	// It is non-nil when DecodedType is DecodedTypeOpenvpn
	Openvpn *decoded.Openvpn `json:"openvpn"`

	// Oracle contains information about an Oracle service.
	// It is non-nil when DecodedType is DecodedTypeOracle
	Oracle *decoded.Oracle `json:"oracle"`

	// PcAnywhere contains information about an PcAnywhere service.
	// It is non-nil when DecodedType is DecodedTypePcAnywhere
	PcAnywhere *decoded.PcAnywhere `json:"pc_anywhere"`

	// Pop3 contains information about an Pop3 service.
	// It is non-nil when DecodedType is DecodedTypePop3
	Pop3 *decoded.Pop3 `json:"pop3"`

	// Postgres contains information about an Postgres service.
	// It is non-nil when DecodedType is DecodedTypePostgres
	Postgres *decoded.Postgres `json:"postgres"`

	// Pptp contains information about an Pptp service.
	// It is non-nil when DecodedType is DecodedTypePptp
	Pptp *decoded.Pptp `json:"pptp"`

	// Prometheus contains information about an Prometheus service.
	// It is non-nil when DecodedType is DecodedTypePrometheus
	Prometheus *decoded.Prometheus `json:"prometheus"`

	// Rdp contains information about an Rdp service.
	// It is non-nil when DecodedType is DecodedTypeRdp
	Rdp *decoded.Rdp `json:"rdp"`

	// Redis contains information about an Redis service.
	// It is non-nil when DecodedType is DecodedTypeRedis
	Redis *decoded.Redis `json:"redis"`

	// S7 contains information about an S7 service.
	// It is non-nil when DecodedType is DecodedTypeS7
	S7 *decoded.S7 `json:"s7"`

	// Sip contains information about an Sip service.
	// It is non-nil when DecodedType is DecodedTypeSip
	Sip *decoded.Sip `json:"sip"`

	// Skinny contains information about an Skinny service.
	// It is non-nil when DecodedType is DecodedTypeSkinny
	Skinny *decoded.Skinny `json:"skinny"`

	// Smtp contains information about an Smtp service.
	// It is non-nil when DecodedType is DecodedTypeSmtp
	Smtp *decoded.Smtp `json:"smtp"`

	// Snmp contains information about an Snmp service.
	// It is non-nil when DecodedType is DecodedTypeSnmp
	Snmp *decoded.Snmp `json:"snmp"`

	// Ssdp contains information about an Ssdp service.
	// It is non-nil when DecodedType is DecodedTypeSsdp
	Ssdp *decoded.Ssdp `json:"ssdp"`

	// Ssh contains information about an Ssh service.
	// It is non-nil when DecodedType is DecodedTypeSsh
	Ssh *decoded.Ssh `json:"ssh"`

	// TeamViewer contains information about an TeamViewer service.
	// It is non-nil when DecodedType is DecodedTypeTeamViewer.
	TeamViewer *decoded.TeamViewer `json:"team_viewer"`

	// Telnet contains information about an Telnet service.
	// It is non-nil when DecodedType is DecodedTypeTelnet
	Telnet *decoded.Telnet `json:"telnet"`

	// Upnp contains information about an Upnp service.
	// It is non-nil when DecodedType is DecodedTypeUpnp
	Upnp *decoded.Upnp `json:"upnp"`

	// Vnc contains information about an Vnc service.
	// It is non-nil when DecodedType is DecodedTypeVnc
	Vnc *decoded.Vnc `json:"vnc"`

	// X11 contains information about an X11 service.
	// It is non-nil when DecodedType is DecodedTypeX11
	X11 *decoded.X11 `json:"x11"`
}
