package main

import (
	"flag"
	"fmt"
	Outlooksend "github.com/svg-rs/SeleniumEmailSender/src"
	"log"
	"os"
	"strings"
	"time"
)

var (
	infoLogger    = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger   = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	var err error
	message, username, password, outlook, gmail, recipientsFile := getFlags()
	var recipients []string
	recipients = handleRecipientsFile(recipientsFile)

	switch {
	case outlook:
		err = Outlooksend.Send(message, username, password, recipients)
		handlerError(err)
	case gmail:
		fmt.Println("Gmail functionality not implemented yet.")
	default:
		help()
	}
}

func handleRecipientsFile(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		handlerError(err)
		return nil
	}

	recipients := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == ',' || r == ';'
	})

	for i := range recipients {
		recipients[i] = strings.TrimSpace(recipients[i])
	}

	return recipients
}

func getFlags() (string, string, string, bool, bool, string) {
	outlook := flag.Bool("outlook", false, "Send email using outlook")
	gmail := flag.Bool("gmail", false, "Send email using gmail")
	message := flag.String("message", "", "Message to send")
	username := flag.String("username", "", "Username to send email")
	password := flag.String("password", "", "Password to send email")
	recipientsFile := flag.String("recipients", "", "Recipients of the email in comma separated format in a text file")

	flag.Parse()

	return *message, *username, *password, *outlook, *gmail, *recipientsFile
}

func handlerError(err error) {
	if err != nil {
		warningLogger.Printf("[ERROR] | %v\n", err)
	}
}

func help() {
	fmt.Println("SeleniumEmailSender")
	fmt.Println("Usage:")
	fmt.Println("  -outlook    Send email using outlook")
	fmt.Println("  -gmail      Send email using gmail")
	fmt.Println("  -message    The message to send")
	fmt.Println("  -username   The username to send email")
	fmt.Println("  -password   The password for the email account")
	fmt.Println("  -recipients  The recipients of the email in comma separated format in a text file")
	fmt.Println("\nFor more information, usage examples, or advanced features, visit the GitHub repository:")
	fmt.Println("  https://github.com/svg-rs/SeleniumEmailSender")
}
