package censys

import (
	"time"
)

// Host represents known data for a Host, obtained from a GetHost call
type Host struct {
	// IP is the IP address of the host.
	IP string `json:"ip"`

	// Name is the Name of the host if it is a Virtual Host, which is provided
	// in the TLS handshake in the Server Name Indication extension
	// (as well as any other place that a domain name could be used at scan time)
	Name string `json:"name"`

	// OperatingSystem is the suspected Operating System of the Host.
	OperatingSystem Software `json:"operating_system"`

	// Location is location of the host.
	Location Location `json:"location"`

	// LocationUpdatedAt is the last time the Location field changed.
	LocationUpdatedAt *time.Time `json:"location_updated_at"`

	// AutonomousSystem countains routing information for the host.
	AutonomousSystem AutonomousSystem `json:"autonomous_system"`

	// AutonomousSystemUpdatedAt is the last time the AutonomousSystem field changed.
	AutonomousSystemUpdatedAt *time.Time `json:"autonomous_system_updated_at"`

	// DNS contains DNS information for the host.
	DNS DNS `json:"dns"`

	// Services is a list of all Services that were observed on the host.
	Services []Service `json:"services"`
}
