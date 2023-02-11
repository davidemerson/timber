package main

import (
    "bufio"
    "fmt"
    "net"
    "time"
)

func main() {
    // Define the length of each interval
    workDuration := time.Second * 30
    restDuration := time.Second * 10

    // Define the number of intervals
    numIntervals := 5

    // Start a new server to listen for incoming connections
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer ln.Close()

    // Keep track of connected clients
    clients := make(map[net.Conn]bool)

    // Start a goroutine to accept incoming connections
    go func() {
        for {
            conn, err := ln.Accept()
            if err != nil {
                fmt.Println("Error accepting connection:", err)
                continue
            }
            clients[conn] = true
            fmt.Println("Accepted new connection from", conn.RemoteAddr())
        }
    }()

    // Start the timer
    for i := 0; i < numIntervals; i++ {
        // Broadcast the start of the work interval to all connected clients
        for conn := range clients {
            _, err := fmt.Fprintln(conn, "Starting work interval")
            if err != nil {
                fmt.Println("Error sending message:", err)
                conn.Close()
                delete(clients, conn)
            }
        }

        time.Sleep(workDuration)

        // Broadcast the start of the rest interval to all connected clients
        for conn := range clients {
            _, err := fmt.Fprintln(conn, "Starting rest interval")
            if err != nil {
                fmt.Println("Error sending message:", err)
                conn.Close()
                delete
