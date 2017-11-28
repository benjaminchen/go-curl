package curl

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode	int
	Headers		map[string]string
	Body		string
}

func NewResponse() *Response {
	return &Response{}
}

func (res *Response) Parse(raw *http.Response) error {
	res.StatusCode = raw.StatusCode
	res.Headers = make(map[string]string)

	for k, v := range raw.Header {
		res.Headers[k] = v[0]
	}

	if body, err := ioutil.ReadAll(raw.Body); err != nil {
		return err
	} else {
		res.Body = string(body)
	}

	return nil
}

func (res *Response) IsOk() bool {
	return res.StatusCode == 200
}