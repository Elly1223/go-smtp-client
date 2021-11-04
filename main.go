package main

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	send("Alarm Test")
}

func send(body string) {
	from := "yourEmail"
	pass := "yourPassword"
	to := "Email"
	port := 587
	hostname := fmt.Sprintf("smtp.office365.com:%d", port)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Alarm Test\n\n" +
		body

	err := smtp.SendMail(hostname,
		LoginAuth(from, pass),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("success!! Please visit your Email Box")
}

type loginAuth struct {
	username, password string
}

// LoginAuth is used for smtp login auth
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}
