package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	w   http.ResponseWriter
	req *http.Request
}

type ApiResponseParams struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	path       string
	Results    any `json:"results"`
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
		path:       r.req.URL.Path,
	})
}

func String(s string) *string {
	return &s
}
