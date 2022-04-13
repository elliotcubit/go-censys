# decoded

This package contains structs describing "decoded" values, or structured data about a particular type of service in Censys data.

For example, consider the following DNS service result from `8.8.8.8`:

```
{
	"_decoded": "dns",
	"dns": {
		"r_code": "SERVER_FAILURE",
		"server_type": "UNKNOWN",
		"resolves_correctly": false
	},
	"extended_service_name": "DNS",
	"observed_at": "2022-04-13T21:59:53.543294317Z",
	"perspective_id": "PERSPECTIVE_TATA",
	"port": 53,
	"service_name": "DNS",
	"source_ip": "167.94.138.44",
	"transport_protocol": "UDP",
	"truncated": false
}
```

The `dns` object is the "decoded".