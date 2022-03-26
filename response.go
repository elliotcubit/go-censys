package censys

import (
	"encoding/json"
	"fmt"
)

type response struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Error  string           `json:"error"`
	Result *json.RawMessage `json:"result"`
	Links  *json.RawMessage `json:"links"`
}

func (r response) asError() error {
	if r.Code == 200 {
		return nil
	}

	return &APIError{
		Code:        r.Code,
		Status:      r.Status,
		ErrorString: r.Error,
	}
}

type APIError struct {
	Code        int
	Status      string
	ErrorString string
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%d: %s; %s", a.Code, a.Status, a.ErrorString)
}
