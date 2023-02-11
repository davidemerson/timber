package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	username := readUsername(conn)
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			break
		}
		fmt.Print(message)
	}
	fmt.Println("Connection closed.")
}

func readUsername(conn net.Conn) string {
	fmt.Print("Enter your username: ")
	username, _ := bufio.NewReader(conn).ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Fprint(conn, username+"\n")
	return username
}
