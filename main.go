package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	defaultServerPort = "8080"
)

var (
	serverPort string
)

func init() {
	serverPort = os.Getenv("SERVER_PORT")
	if len(serverPort) == 0 {
		serverPort = defaultServerPort
	}
}

func main() {
	http.HandleFunc("/", httpHandler)
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		panic(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Proto: %s\n", r.Proto)
	fmt.Printf("RemoteAddress: %s\n", r.RemoteAddr)
	fmt.Printf("Request URL: %s\n", r.URL)
	fmt.Println("Request Header:")
	for k, v := range r.Header {
		fmt.Printf("  » %-20s%s\n", k, v)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	fmt.Printf("Request Body: \n---\n%v\n---\n", string(body))

	_, err = w.Write([]byte("success"))
	if err != nil {
		panic(err)
	}
}
