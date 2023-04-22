package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    http.HandleFunc("/", handleTimer)
    fmt.Println("Starting server on port 8080...")
    http.ListenAndServe(":8080", nil)
}

func handleTimer(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("config.txt")
    if err != nil {
        fmt.Fprintf(w, "Error: %v", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var duration int
    var interval int
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "=")
        if len(parts) != 2 {
            continue
        }
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        if key == "duration" {
            duration, _ = strconv.Atoi(value)
        } else if key == "interval" {
            interval, _ = strconv.Atoi(value)
        }
    }

    ticker := time.NewTicker(time.Duration(interval) * time.Second)
    startTime := time.Now()
    endTime := startTime.Add(time.Duration(duration) * time.Second)
    for now := range ticker.C {
        remaining := endTime.Sub(now)
        if remaining <= 0 {
            fmt.Fprintf(w, "Time's up!")
            ticker.Stop()
            return
        }
        fmt.Fprintf(w, "%v\n", remaining)
    }
}
