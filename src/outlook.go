package outlook

import (
	"log"
	"os"
)

var (
	infoLogger    = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger   = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func outlook(message, username, password string) {
	if message == "" || username == "" || password == "" {
		errorLogger.Println("Message or username or password is missing!")
	}

	infoLogger.Println("%v, %v, %v", username, password, message)

}
