package http_util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, header []Header) (result string) {
	fmt.Print("url is " + url)
	req, _ := http.NewRequest("GET", url, nil)
	setHttpHeader(header, req)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	return string(body)
}

func setHttpHeader(params []Header, request *http.Request) {
	for _, v := range params {
		request.Header.Add(v.Key, v.Value)
	}
}
