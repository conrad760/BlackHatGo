package portscan

import (
	"fmt"
	"net"
	"sort"
)

// scans a single port given a url:portnumber
func scanme(t string, url string) (res string) {
	_, err := net.Dial(t, url)
	if err == nil {
		fmt.Println("Connection Success")
		res = "Connection Success"
	}
	return res
}
func scan1024() (res []int) {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 81; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 81; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
	return openports
}

func worker(ports, results chan int) {
	url := "scanme.nmap.org"
	t := "tcp"
	for p := range ports {
		address := fmt.Sprintf("%s:%d", url, p)
		conn, err := net.Dial(t, address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	scan1024()
}
