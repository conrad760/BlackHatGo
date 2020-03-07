package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data.")
	}
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Wait for connection
		conn, err := listener.Accept()
		log.Println("recieved connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using gorouting for concurrency.
		go echo(conn)
	}

}
