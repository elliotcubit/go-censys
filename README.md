# go-censys

`go-censys` is an Go API wrapper around the [Censys Search API](https://search.censys.io/api).

This is *_NOT_* and official Censys library, and it is not supported by or affiliated with Censys at this time. I do not own Censys trademarks or copyrights.

## Installation

Add `go-censys` to your project via [Go Modules](https://go.dev/ref/mod):

```
go get github.com/elliotcubit/go-censys
```

## Quickstart

You will need a Censys API ID and Secret to use `go-censys`. After logging into a Censys account, they can be retrieved on the [account details page](https://search.censys.io/account/api).

This example quickstart fetches all services on `209.216.230.240` from Censys search, including services that are hosted behind SNI, and prints them plus the hash of the certificate they serve (if any).

```
package main

import (
	"fmt"
	"os"

	"github.com/elliotcubit/go-censys"
)

func main() {
	id := os.Getenv("CENSYS_ID")
	secret := os.Getenv("CENSYS_SECRET")

	// The third (nil) argument here is an optional *http.Client to use for requests.
	client := censys.NewClient(id, secret, nil)

	// The last (empty string) argument here is a token for fetching results
	// beyond the first page using the token in the returned `results`.
	results, err := client.SearchHosts("ip: 209.216.230.240", 100, censys.VirtualHostsInclude, "")
	if err != nil {
		fmt.Printf("Error fetching search results: %s\n", err.Error())
		return
	}

	for _, host := range results.Hits {
		if len(host.Name) > 0 {
			fmt.Printf("%s (%s)\n", host.IP, host.Name)
		} else {
			fmt.Printf("%s\n", host.IP)
		}

		for _, service := range host.Services {
			fmt.Printf(
				"\t%s %d: %s",
				service.TransportProtocol,
				service.Port,
				service.ServiceName,
			)

			if len(service.Certificate) > 0 {
				fmt.Printf(" (%s)\n", service.Certificate)
			} else {
				fmt.Println("")
			}
		}

		fmt.Println("")
	}
}
```

Example output:

```
209.216.230.240
        TCP 80: HTTP
        TCP 443: HTTP (34556ed249da9c079c6ab4394a7ec0faeb58802ade1131a21952655e2f9c601e)

209.216.230.240 (news.ycombinator.com)
        TCP 80: HTTP
        TCP 443: HTTP (34556ed249da9c079c6ab4394a7ec0faeb58802ade1131a21952655e2f9c601e)
```

## Supported Endpoints

Note that some endpoints may be listed twice (in separate categories) to match the [official](https://search.censys.io/api) API documentation.

### Hosts

- [x] `GET /v2/hosts/search`
- [x] `GET /v2/hosts/aggregate`
- [x] `GET /v2/hosts/{ip}`
- [ ] `GET /v2/hosts/{ip}/diff`
- [ ] `GET /v2/experimental/hosts/{ip}/events`
- [x] `GET /v2/hosts/{ip}/names`
- [ ] `GET /v2/hosts/{ip}/comments`
- [ ] `POST /v2/hosts/{ip}/comments`
- [ ] `GET /v2/hosts/{ip}/comments/{comment_id}`
- [ ] `PUT /v2/hosts/{ip}/comments/{comment_id}`
- [ ] `DELETE /v2/hosts/{ip}/comments/{comment_id}`
- [x] `GET /v2/metadata/hosts`
- [ ] `GET /v2/tags/{id}/hosts`
- [ ] `GET /v2/hosts/{ip}/tags`
- [ ] `PUT /v2/hosts/{ip}/tags/{id}`
- [ ] `DELETE /v2/hosts/{ip}/tags/{id}`

### Certificates V1

- [ ] `GET /v1/view/certificates/{sha256}`
- [ ] `POST /v1/search/certificates`
- [ ] `POST /v1/report/certificates`
- [ ] `POST /v1/bulk/certificates`

### Certificates V2

- [ ] `GET /v2/certificates/{fingerprint}/hosts`
- [ ] `GET /v2/certificates/{fingerprint}/comments`
- [ ] `POST /v2/certificates/{fingerprint}/comments`
- [ ] `GET /v2/certificates/{fingerprint}/comments/{comment_id}`
- [ ] `PUT /v2/certificates/{fingerprint}/comments/{comment_id}`
- [ ] `DELETE /v2/certificates/{fingerprint}/comments/{comment_id}`
- [ ] `GET /v2/tags/{id}/certificates`
- [ ] `GET /v2/certificates/{fingerprint}/tags`
- [ ] `PUT /v2/certificates/{fingerprint}/tags/{id}`
- [ ] `DELETE /v2/certificates/{fingerprint}/tags/{id}`

### Account

- [ ] `GET /v1/account`

### Metadata

- [x] `GET /v2/metadata/hosts`

### Comments

#### On hosts...
- [ ] `GET /v2/hosts/{ip}/comments`
- [ ] `POST /v2/hosts/{ip}/comments`
- [ ] `GET /v2/hosts/{ip}/comments/{comment_id}`
- [ ] `PUT /v2/hosts/{ip}/comments/{comment_id}`
- [ ] `DELETE /v2/hosts/{ip}/comments/{comment_id}`

#### On certificates...
- [ ] `GET /v2/certificates/{fingerprint}/comments`
- [ ] `POST /v2/certificates/{fingerprint}/comments`
- [ ] `GET /v2/certificates/{fingerprint}/comments/{comment_id}`
- [ ] `PUT /v2/certificates/{fingerprint}/comments/{comment_id}`
- [ ] `DELETE /v2/certificates/{fingerprint}/comments/{comment_id}`

### Tags

- [ ] `GET /v2/tags`
- [ ] `POST /v2/tags`
- [ ] `GET /v2/tags/{id}`
- [ ] `PUT /v2/tags/{id}`
- [ ] `DELETE /v2/tags/{id}`

#### On hosts...
- [ ] `GET /v2/hosts/{ip}/tags`
- [ ] `PUT /v2/hosts/{ip}/tags/{id}`
- [ ] `DELETE /v2/hosts/{ip}/tags/{id}`

#### On certificates...

- [ ] `GET /v2/certificates/{fingerprint}/tags`
- [ ] `PUT /v2/certificates/{fingerprint}/tags/{id}`
- [ ] `DELETE /v2/certificates/{fingerprint}/tags/{id}`

### Experimental

- [ ] `GET /v2/experimental/hosts/{ip}/events`

## License

MIT