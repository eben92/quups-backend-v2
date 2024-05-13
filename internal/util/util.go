package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
}

type ApiResponseParams struct {
	StatusCode int     `json:"status_code"`
	Message    *string `json:"message"`
	path       *string
	Results    any `json:"results"`
}

func (r *Response) Builder(w http.ResponseWriter, req *http.Request, data *ApiResponseParams) ([]byte, error) {
	w.WriteHeader(data.StatusCode)

	return json.Marshal(&ApiResponseParams{
		Results:    data.Results,
		StatusCode: data.StatusCode,
		Message:    data.Message,
		path:       &req.URL.Path,
	})
}

func String(s string) *string {
	return &s
}
