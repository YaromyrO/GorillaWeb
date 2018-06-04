package client

import (
	"net/http"
	"fmt"
	"os"
	"bytes"
	"io/ioutil"
)

func Listener() {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	for _, arg := range catch() {
		buffer.WriteString(arg)
	}
	doRequest(buffer.String())
}

func doRequest(url string) {
	request, err := http.NewRequest("GET", url, nil)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Add("cache-control", "no-cache")

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Body)
}

func catch() []string {
	return os.Args[1:]
}