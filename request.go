package curl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"io"
	"errors"
	"net/http"
)

type Request struct {
	Url			string
	Headers		map[string]string
	Cookies		map[string]string
	Queries		map[string]string
	PostData	map[string]interface{}
	Timeout  time.Duration
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.Headers = headers
	return r
}

func (r *Request) SetQueries(queries map[string]string) *Request {
	r.Queries = queries
	return r
}

func (r *Request) SetCookies(cookies map[string]string) *Request {
	r.Cookies = cookies
	return r
}

func (r *Request) SetPostData(postData map[string]interface{}) *Request {
	r.PostData = postData
	return r
}

func (r *Request) SetTimeout(timeout time.Duration) *Request {
	r.Timeout = timeout
	return r
}

func (r *Request) Get() (*Response, error) {
	return r.send(http.MethodGet)
}

func (r *Request) Post() (*Response, error) {
	return r.send(http.MethodPost)
}

func (r *Request) send(method string) (response *Response, err error) {
	response = NewResponse()

	if r.Url == "" {
		err = errors.New("empty request url, please set url first")
		return
	}

	var payload io.Reader

	// set post data
	if method == http.MethodPost && len(r.PostData) > 0 {
		var data []byte
		if data, err = json.Marshal(r.PostData); err != nil {
			return
		} else {
			payload = bytes.NewReader(data)
		}
	} else {
		payload = nil
	}

	var req *http.Request
	if req, err = http.NewRequest(method, r.Url, payload); err != nil {
		return
	}

	// set headers
	if len(r.Headers) > 0 {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}

	// set queries
	if len(r.Queries) > 0 {
		q := req.URL.Query()
		for k, v := range r.Queries {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	// set cookies
	if len(r.Cookies) > 0 {
		for k, v := range r.Cookies {
			req.AddCookie(&http.Cookie{
				Name: k,
				Value: v,
			})
		}
	}

	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		return
	} else {
		response.Parse(res)
	}

	// clear body and close
	defer func() {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	return
}
