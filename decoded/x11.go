package decoded

type X11 struct {
	RefusalReason          string `json:"refusal_reason"`
	RequiresAuthentication bool   `json:"requires_authentication"`
	Vendor                 string `json:"vendor"`
	Version                string `json:"version"`
}
