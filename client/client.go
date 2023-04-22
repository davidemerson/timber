package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	config, err := loadConfig("config.txt")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	client := &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}

	resp, err := client.Get(config.ServerURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", body)
}

type clientConfig struct {
	ServerURL string
	Timeout   int
}

func loadConfig(filename string) (*clientConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	config := &clientConfig{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if key == "server_url" {
			config.ServerURL = value
		} else if key == "timeout" {
			config.Timeout, _ = strconv.Atoi(value)
		}
	}

	return config, nil
}
