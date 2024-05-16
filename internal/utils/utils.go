package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"strings"
)

type Response struct {
	w   http.ResponseWriter
	req *http.Request
}

type ApiResponseParams struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Path       string `json:"path"`
	Results    any    `json:"results"`
}

const (
	MSISDN_PREFIX_P233    string = "+233"
	DEFAULT_MSISDN_PREFIX string = "233"
)

func New(w http.ResponseWriter, r *http.Request) *Response {

	return &Response{
		w:   w,
		req: r,
	}

}

func (r *Response) WrapInApiResponse(data *ApiResponseParams) ([]byte, error) {
	r.w.WriteHeader(data.StatusCode)

	return json.Marshal(&ApiResponseParams{
		Results:    data.Results,
		StatusCode: data.StatusCode,
		Message:    data.Message,
		Path:       r.req.URL.Path,
	})
}

func String(s string) *string {
	return &s
}

// ReplacePrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func ReplacePrefix(s, prefix, with string) string {

	if strings.HasPrefix(s, prefix) {
		return with + strings.TrimPrefix(s, prefix)
	}

	return s

}

// eg. 0550404071 will be converted to 233550404071 in (bytes)
//
//	If msisdn doesn't start with prefix, msisdn is returned unchanged.
func ConvertToLocalMsisdn(msisdn string) ([]byte, error) {

	if len(msisdn) < 9 || len(msisdn) > 13 {
		return nil, fmt.Errorf("invalid phone number")
	}

	msisdn = ReplacePrefix(msisdn, "0", DEFAULT_MSISDN_PREFIX)
	msisdn = ReplacePrefix(msisdn, MSISDN_PREFIX_P233, DEFAULT_MSISDN_PREFIX)

	if !strings.HasPrefix(msisdn, DEFAULT_MSISDN_PREFIX) {
		return []byte(msisdn), errors.New("Invalid phone number")
	}

	return []byte(msisdn), nil
}

// IsValidMsisdn removes any local prefix in msisdn(+233 or 0), returns
// the updated msisdn(233xxxx...) and a boolean value
//   - true (if msisdn is not an empty string) otherwise false.
//     If msisdn doesn't start with prefix, msisdn is returned unchanged.
func IsValidMsisdn(msisdn string) (string, bool) {

	m, err := ConvertToLocalMsisdn(msisdn)

	if err != nil {
		return string(m), false
	}

	return string(m), true
}

func IsVaildEmail(e string) bool {

	_, err := mail.ParseAddress(e)

	return err == nil
}
