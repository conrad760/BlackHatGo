package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	resp, err := http.Get("https://conrad-parker.com/")
	if err != nil {
		log.Panicln(err)
	}
	// Print HTTP Status
	fmt.Println("1)", resp.Status)

	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("2)", string(body))
	resp.Body.Close()

	resp, err = http.Head("https://conrad-parker.com")
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println("3)", resp.Status)

	form := url.Values{}
	form.Add("foo", "bar")
	resp, err = http.Post(
		"https://conrad-parker.com",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println("4)", resp.Status)

	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err = client.Do(req)
	resp.Body.Close()
	fmt.Println("5)", resp.Status)

	req, err = http.NewRequest("PUT", "https://www.google.com/robots.txt", strings.NewReader(form.Encode()))
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println("6)", resp.Status)
}
