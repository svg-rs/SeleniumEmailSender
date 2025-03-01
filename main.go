package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	var err error
	err = godotenv.Load(".env")
	handlerError(err)

	var message string = os.Getenv("MESSAGE")
	var username string = os.Getenv("USERNAME")
	var password string = os.Getenv("PASSWORD")

	outlook.
}

func handlerError(err error) {
	if err != nil {
		log.Printf("[ERROR] | %v\n", err)
	} else {
		return
	}
}
