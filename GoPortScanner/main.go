package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var host string
var startPort int
var endPort int

func main() {
	userInput()
}

func checkPort(host string, startPort, endPort int) {

	for i := startPort; i <= endPort; i++ {
		fmt.Println(i)
		qualifiedHost := fmt.Sprintf("%s%s%d", host, ":", i)
		conn, err := net.DialTimeout("tcp", qualifiedHost, 1*time.Second)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status)
	}
}

func userInput() {
	fmt.Println("Host> ")
	fmt.Scan(&host)
	fmt.Println("Starting Port (i. e. 80)> ")
	fmt.Scan(&startPort)
	fmt.Println("End Port (i. e. 8080)> ")
	fmt.Scan(&endPort)
	fmt.Println("Running scan...")
	checkPort(host, startPort, endPort)
}
