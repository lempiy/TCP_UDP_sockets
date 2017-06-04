package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	connTCPAddr, _ := net.ResolveTCPAddr("tcp4", ":1200")
	connSocket, err := net.DialTCP("tcp", nil, connTCPAddr)
	if err != nil {
		fmt.Println(err)
	}
	_, err = connSocket.Write([]byte("any"))
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(connSocket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
}
