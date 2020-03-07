// +build unit

package portscan

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CheckConnection struct {
	t        string
	url      string
	expected string
}

var connectingResults = []CheckConnection{
	{"tcp", "scanme.nmap.org:80", "Connection Success"},
	{"tcp", "google.com:80", "Connection Success"},
	{"tcp", "conrad-parker.com:80", "Connection Success"},
	{"tcp", "reamer.house:80", "Connection Success"},
}

func TestForOpenPorts(t *testing.T) {
	results := scan1024()
	if len(results) < 0 {
		t.Fatal("Expected open ports")
	}
}

func TestConnectionSuccess(t *testing.T) {
	for _, test := range connectingResults {
		result := scanme(test.t, test.url)
		if result != test.expected {
			t.Fatal("Test Failed: {} and {} inputted, {} expected, received: {}", test.t, test.url, test.expected, result)
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world" {
		t.Fatal("String contents do not match expected")
	}
}

func TestHTTPRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"200\" }")
	}
	req := httptest.NewRequest("GET", "https://conrad-parker.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code not 200")
	}
}
