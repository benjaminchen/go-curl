package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Request stores settings.
type Request struct {
	URL      string
	Headers  map[string]string
	Cookies  []*http.Cookie
	Queries  map[string]string
	PostData map[string]interface{}
	Timeout  time.Duration
}

// NewRequest creates a Request instance.
func NewRequest() *Request {
	return &Request{}
}

// SetURL sets request url.
func (r *Request) SetURL(url string) *Request {
	r.URL = url
	return r
}

// SetHeaders sets request headers.
func (r *Request) SetHeaders(headers map[string]string) *Request {
	r.Headers = headers
	return r
}

// SetQueries sets request queries.
func (r *Request) SetQueries(queries map[string]string) *Request {
	r.Queries = queries
	return r
}

// SetCookies sets request cookies.
func (r *Request) SetCookies(cookies []*http.Cookie) *Request {
	r.Cookies = cookies
	return r
}

// SetPostData sets request post datas.
func (r *Request) SetPostData(postData map[string]interface{}) *Request {
	r.PostData = postData
	return r
}

// SetTimeout sets request timeout.
func (r *Request) SetTimeout(timeout time.Duration) *Request {
	r.Timeout = timeout
	return r
}

// Get uses request get method
func (r *Request) Get() (*Response, error) {
	return r.send(http.MethodGet)
}

// Post uses request post method
func (r *Request) Post() (*Response, error) {
	return r.send(http.MethodPost)
}

func (r *Request) send(method string) (response *Response, err error) {
	response = NewResponse()

	if r.URL == "" {
		err = errors.New("empty request url, please set url first")
		return
	}

	var payload io.Reader

	// set post data
	if method == http.MethodPost && len(r.PostData) > 0 {
		var data []byte
		if data, err = json.Marshal(r.PostData); err != nil {
			return
		}

		payload = bytes.NewReader(data)
	} else {
		payload = nil
	}

	var req *http.Request
	if req, err = http.NewRequest(method, r.URL, payload); err != nil {
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
		for _, v := range r.Cookies {
			req.AddCookie(v)
		}
	}

	var res *http.Response
	var client *http.Client
	if r.Timeout > 0 {
		client = &http.Client{
			Timeout: r.Timeout,
		}
	}

	if res, err = client.Do(req); err != nil {
		return
	}

	response.Parse(res)

	// clear body and close
	defer func() {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	return
}
