package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set auth method
	auth := smtp.PlainAuth("script@localhost", "script", "password", "localhost")

	// Connect to server, set up message and send it
	to := []string{"admin@localhost"}
	msg := []byte("To: admin@localhost\n" +
		"Subject: Test email.\n" +
		"\n" +
		"Test message.\n")
	err := smtp.SendMail("localhost:25", auth, "script@localhost", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
