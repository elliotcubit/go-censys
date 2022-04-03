package censys

type HostAggregate struct {
	Total              int                   `json:"total"`
	TotalOmitted       int                   `json:"total_omitted"`
	PotentialDeviation int                   `json:"potential_deviation"`
	Buckets            []HostAggregateBucket `json:"buckets"`
	Query              string                `json:"query"`
	Field              string                `json:"field"`
}

type HostAggregateBucket struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
}
