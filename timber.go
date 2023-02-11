package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "time"
)

const (
    workInterval = 20 * time.Second
    restInterval = 10 * time.Second
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server started on port 8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    username := readUsername(conn)
    fmt.Println("Accepted connection from", username)
    startTimer(conn, username)
}

func readUsername(conn net.Conn) string {
    reader := bufio.NewReader(conn)
    fmt.Fprint(conn, "Enter your username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)
    return username
}

func startTimer(conn net.Conn, username string) {
    startTime := time.Now().Add(5 * time.Second)
    timer := time.NewTimer(startTime.Sub(time.Now()))
    for i := 0; i < 5; i++ {
        <-timer.C
        fmt.Fprintln(conn, username, "Work interval started...")
        time.Sleep(workInterval)
        fmt.Fprintln(conn, username, "Work interval ended.")
        fmt.Fprintln(conn, username, "Rest interval started...")
        time.Sleep(restInterval)
        fmt.Fprintln(conn, username, "Rest interval ended.")
        timer.Reset(workInterval + restInterval)
    }
    fmt.Fprintln(conn, username, "Timer ended.")
}
