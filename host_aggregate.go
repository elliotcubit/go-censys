package censys

type Aggregate struct {
	Total              float64           `json:"total"`
	TotalOmitted       float64           `json:"total_omitted"`
	PotentialDeviation float64           `json:"potential_deviation"`
	Buckets            []AggregateBucket `json:"buckets"`
	Query              string            `json:"query"`
	Field              string            `json:"field"`
}

type AggregateBucket struct {
	Key   string  `json:"key"`
	Count float64 `json:"count"`
}
