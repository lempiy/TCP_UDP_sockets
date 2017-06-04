package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maximum request length to 128B to prevent flood based attacks
	defer conn.Close()                                    // close connection before exit
	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if readLen == 0 {
			break // connection already closed by client
		} else if string(request[:readLen]) == "timestamp" {
			fmt.Println(string(request[:readLen]))
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
			break
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
			break
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
