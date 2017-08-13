package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var host string
var start_port int
var end_port int

func main() {
	user_input()
}

func check_port(host string, start_port, end_port int) {

	for i := start_port; i <= end_port; i++ {
		//fmt.Println(i)
		qualified_host := fmt.Sprintf("%s%s%d", host, ":", i) //host, duepunti porta incrementale
		conn, err := net.DialTimeout("tcp", qualified_host, 1*time.Second)
		if err != nil {
			//fmt.Println(err)
			continue
		}
		fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("porta: ", i, " - ", "status: ", status)
	}
}

func user_input() {
	fmt.Println("Host> ")
	fmt.Scan(&host)
	fmt.Println("Starting Port (i. e. 80)> ")
	fmt.Scan(&start_port)
	fmt.Println("End Port (i. e. 8080)> ")
	fmt.Scan(&end_port)
	fmt.Println("Running scan...")
	check_port(host, start_port, end_port)
}
