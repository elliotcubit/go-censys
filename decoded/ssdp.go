package decoded

type Ssdp struct {
	Headers HttpHeaders `json:"headers"`
	UpnpUrl string      `json:"upnp_url"`
}
