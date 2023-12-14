package main

import (
	"os"
	"log"
	"page-hit-tracker/pkg/router"
)

func main() {
	if err := router.Run(); err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}
