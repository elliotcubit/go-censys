package decoded

type Modbus struct {
	ExceptionResponse Modbus_ExceptionResponse `json:"exception_response"`
	Function          uint64                   `json:"function"`
	MeiResponse       Modbus_MeiResponse       `json:"mei_response"`
	UnitId            int64                    `json:"unit_id"`
}

type Modbus_MeiResponse struct {
	ConformityLevel int64                       `json:"conformity_level"`
	MoreFollows     bool                        `json:"more_follows"`
	Objects         []Modbus_MeiResponse_Object `json:"objects"`
}

type Modbus_MeiResponse_Object struct {
	ModelName           string `json:"model_name"`
	ProductCode         string `json:"product_code"`
	ProductName         string `json:"product_name"`
	Revision            string `json:"revision"`
	UserApplicationName string `json:"user_application_name"`
	Vendor              string `json:"vendor"`
	VendorUrl           string `json:"vendor_url"`
}

type Modbus_ExceptionResponse struct {
	ExceptionFunction uint64 `json:"exception_function"`
	ExceptionType     uint64 `json:"exception_type"`
}
