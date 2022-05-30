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
	results, err := client.SearchHosts(
		"ip: 209.216.230.240",
		100,
		censys.VirtualHostsInclude,
		censys.SortRelevance,
		"",
	)
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
