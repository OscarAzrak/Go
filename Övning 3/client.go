package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Add a time limit for all requests made by this client.
	client := &http.Client{Timeout: 10 * time.Second}

	for {
		before := time.Now()
		//res := Get(server[0], client)
		res := MultiGet(server, client)
		after := time.Now()
		fmt.Println("Response:", res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
}

type Response struct {
	Body       string
	StatusCode int
}

func (r *Response) String() string {
	return fmt.Sprintf("%q (%d)", r.Body, r.StatusCode)
}

// Get makes an HTTP Get request and returns an abbreviated response.
// The response is empty if the request fails.
func Get(url string, client *http.Client) *Response {
	res, err := client.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), res.StatusCode}
}

// MultiGet makes an HTTP Get request to each url and returns
// the response from the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is 503
// – Service unavailable.
func MultiGet(urls []string, client *http.Client) *Response {

    before := time.Now()
    timeout := time.After(700*time.Millisecond)

    Fch := make(chan*Response,1)
    resp := make(chan*Response,1)

    for _, url := range(urls) {
        go func(url string) {
            response := Get(url, client)
            if response.StatusCode == 503 {
                Fres := &Response{"Service unavailable", 503}
				Fch <- Fres
            }
            if response.StatusCode == 200 {
				resp <- response
			}
        }(url)
    }
	select {

	case timer := <-timeout:
		fmt.Println("Time: ", timer.Sub(before))
		return &Response{"Timeout error", 503}
	case response := <-resp:
		return response
	case Fres := <-Fch:
		return Fres
    }
}