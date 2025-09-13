package main

import (
	"log"
	"vmopt-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal("Error: %v", err)
	}
}
