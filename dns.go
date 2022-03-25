package censys

import (
	"time"
)

type DNS struct {
	Names   []string             `json:"name"`
	Records map[string]DNSRecord `json:"records"`
}

type DNSRecord struct {
	Recordtype string    `json:"record_type"`
	ResolvedAt time.Time `json:"resolved_at"`
}

type ReverseDNS struct {
	Names      []string  `json:"names"`
	ResolvedAt time.Time `json:"resolved_at"`
}
