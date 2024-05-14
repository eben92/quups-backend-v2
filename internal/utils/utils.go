package utils

import (
	"encoding/json"
	"net/http"
	"net/mail"
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

func ValidateMsisdn(s string) {
	// TODO
}

func IsVaildEmail(e string) bool {

	_, err := mail.ParseAddress(e)

	return err == nil
}
