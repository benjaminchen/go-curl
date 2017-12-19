package curl

import (
	"io/ioutil"
	"net/http"
)

// Response stores http response datas.
type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
	Raw        *http.Response
}

// NewResponse creates Response instance.
func NewResponse() *Response {
	return &Response{}
}

// Parse parses http response.
func (res *Response) Parse(raw *http.Response) (err error) {
	res.Raw = raw
	res.StatusCode = raw.StatusCode
	res.Headers = make(map[string]string)

	for k, v := range raw.Header {
		res.Headers[k] = v[0]
	}

	var body []byte
	if body, err = ioutil.ReadAll(raw.Body); err != nil {
		return
	}

	res.Body = string(body)

	return nil
}

// IsOk checks http request ok or not
func (res *Response) IsOk() bool {
	return res.StatusCode == 200
}
