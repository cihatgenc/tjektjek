package main

import (
	"fmt"
	"log"
	"os"
)

// GetHostName - Get hostname
func GetHostName() (string, error) {
	fmt.Printf("Executing GetHostName\n")
	name, err := os.Hostname()

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return name, nil
}
