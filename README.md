# GO cURL library

## Install

```
go get github.com/benjaminchen/go-curl
```

## Usage

### Get

```
package main

import (
	"fmt"
	"github.com/benjaminchen/go-curl"
)

func main() {
	url := "your-url"

	headers := map[string]string{
		"Content-Type":  "application/json",
	}

	cookies := map[string]string {
		"yourCookieKey": "yourCookieValue",
	}

	queries := map[string]string {
		"yourQueryKey": "phyourQueryValue",
	}

	curl := curl.NewRequest()
	response, err := curl.SetUrl(url).SetHeaders(headers).SetCookies(cookies).SetQueries(queries).Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Headers)
		fmt.Println(response.Body)
	}
}
```

### Post

```
package main

import (
	"fmt"
	"github.com/benjaminchen/go-curl"
)

func main() {
	url := "your-url"

	headers := map[string]string{
		"Content-Type":  "application/json",
	}

	cookies := map[string]string {
		"yourCookieKey": "yourCookieValue",
	}

	queries := map[string]string {
		"yourQueryKey": "phyourQueryValue",
	}

	postData := map[string]interface{} {
		"yourPostData": "data",
	}

	curl := curl.NewRequest()
	response, err := curl.SetUrl(url).SetHeaders(headers).SetCookies(cookies).SetQueries(queries).SetPostData(postData).Post()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Headers)
		fmt.Println(response.Body)
	}
}
```