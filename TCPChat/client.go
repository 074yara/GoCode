package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8008")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go clientReader(conn)
	output := bufio.NewScanner(os.Stdin)
	for output.Scan() {
		_, err = conn.Write([]byte(output.Text() + "\n"))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}

func clientReader(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
