package main

import (
	"./myserver"
	"net/http"
	"os"
	"log"
)

func main() {
	r := myserver.MakeRouter()
	os.Setenv("Address", "localhost:8080")
	srv := &http.Server{
		Handler: r,
		Addr:    os.Getenv("Address"),
	}
	log.Fatal(srv.ListenAndServe())
}

//func listener() {
//	var buffer bytes.Buffer
//	buffer.WriteString("http://")
//	for _, arg := range catch() {
//		buffer.WriteString(arg)
//	}
//	doRequest(buffer.String())
//}
//
//func doRequest(url string) {
//	request, err := http.NewRequest("GET", url, nil)
//	response, err := http.DefaultClient.Do(request)
//	if err != nil {
//		fmt.Println(err)
//	}
//	request.Header.Add("cache-control", "no-cache")
//
//	defer response.Body.Close()
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(body))
//}

func catch() []string {
	return os.Args[1:]
}
