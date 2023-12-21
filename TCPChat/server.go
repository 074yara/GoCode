package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	mu       sync.Mutex
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case user := <-leaving:
			mu.Lock()
			delete(clients, user)
			mu.Unlock()

		case user := <-entering:
			mu.Lock()
			clients[user] = true
			mu.Unlock()
		case msg := <-messages:
			mu.Lock()
			for eachClient := range clients {
				if clients[eachClient] == true {
					eachClient.ch <- msg
				}
			}
			mu.Unlock()
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func connHandler(conn net.Conn) {
	who := conn.RemoteAddr().String()
	ch := make(chan string)
	user := client{
		conn: conn,
		name: who,
		ch:   ch,
	}
	go clientWriter(conn, ch)
	messages <- who + " has arrived"
	entering <- user
	ch <- "you are: " + who
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- user
	messages <- who + " has left"
	_ = conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8008")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go connHandler(conn)
	}
}
