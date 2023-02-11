package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strconv"
    "strings"
    "time"
)

type Config struct {
    port     int
    duration int
    interval int
}

func main() {
    config := loadConfig("config.txt")
    listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.port))
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening on port", config.port)
    clients := make(map[net.Conn]string)
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn, clients, config.duration, config.interval)
    }
}

func loadConfig(filename string) Config {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening config file:", err)
        os.Exit(1)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var config Config
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "=")
        key := parts[0]
        value := parts[1]
        switch key {
        case "port":
            port, err := strconv.Atoi(value)
            if err != nil {
                fmt.Println("Error parsing port:", err)
                os.Exit(1)
            }
            config.port = port
        case "duration":
            duration, err := strconv.Atoi(value)
            if err != nil {
                fmt.Println("Error parsing duration:", err)
                os.Exit(1)
            }
            config.duration = duration
        case "interval":
            interval, err := strconv.Atoi(value)
            if err != nil {
                fmt.Println("Error parsing interval:", err)
                os.Exit(1)
            }
            config.interval = interval
        }
    }
    return config
}

func readUsername(conn net.Conn) string {
    username, _ := bufio.NewReader(conn).ReadString('\n')
    username = strings.TrimSpace(username)
    return username
}

func handleConnection(conn net.Conn, clients map[net.Conn]string, duration int, interval int) {
    username := readUsername(conn)
    clients[conn] = username
    fmt.Println("Accepted connection from", username)
    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("Error reading from", username+":", err)
            delete(clients, conn)
            conn.Close()
            break
        }
        if message == "start\n" {
            fmt.Println(username, "has started the timer")
            start := time.Now()
            for i := 0; i < duration/interval; i++ {
                time.Sleep(time.Duration(interval) * time.Second)
                for client := range clients {
                    _, err := client.Write([]byte(strconv.Itoa(interval*(i+1)) + "\n"))
                    if err != nil {
                        fmt.Println("Error writing to", clients[client]+":", err)
                        client.Close()
                        delete(clients, client)
                    }
                }
            }
            fmt.Println(username, "timer has completed in", time.Since(start))
        } else {
            fmt.Println("Received message from", username+":", message)
        }
    }
}
