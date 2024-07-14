package local_http

import (
	"io"
	"net/http"
)

type Options struct {
	Method  *string
	Body    io.Reader
	Headers *[][2]string
}

// Fetch is a simple method to handle simple http fetch operations
//
// eg.
//
//		res, err := local_http.Fetch("https://api.com/", &local_http.Options{
//			Method: "GET", // default method: GET
//			Body: nil,
//			Headers: &[][2]string{
//				{"Authorization": "Bearer ey20920"}
//			}
//		})
//
//	 if err != nil {
//			// handle error
//		}
//		defer res.Body.Close()
//
//	 // decode res.Body
func Fetch(url string, options *Options) (*http.Response, error) {
	client := &http.Client{}

	if options == nil {
		options = &Options{}
	}

	if options.Method == nil {
		GET := http.MethodGet
		options.Method = &GET
	}

	if options.Headers == nil {
		options.Headers = &[][2]string{}
	}

	req, err := http.NewRequest(*options.Method, url, options.Body)

	headers := *options.Headers

	if len(headers) > 0 {
		for _, value := range headers {
			req.Header.Set(value[0], value[1])

		}
	}

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
