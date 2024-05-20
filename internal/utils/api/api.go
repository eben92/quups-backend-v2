package apiutils

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
	Path       string `json:"path"`
	Results    any    `json:"results"`
}

func New(w http.ResponseWriter, r *http.Request) *Response {

	return &Response{
		w:   w,
		req: r,
	}

}

func (r *Response) WrapInApiResponse(data *ApiResponseParams) {
	r.w.WriteHeader(data.StatusCode)

	res, _ := json.Marshal(&ApiResponseParams{
		Results:    data.Results,
		StatusCode: data.StatusCode,
		Message:    data.Message,
		Path:       r.req.URL.Path,
	})

	_, _ = r.w.Write(res)
}
