package decoded

type Telnet struct {
	Banner string            `json:"banner"`
	Do     map[uint64]string `json:"do"`
	Dont   map[uint64]string `json:"dont"`
	Will   map[uint64]string `json:"will"`
	Wont   map[uint64]string `json:"wont"`
}
