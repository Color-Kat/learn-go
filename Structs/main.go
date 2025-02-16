package main

import "fmt"

type user struct {
	name  string
	email string
	phone string
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

func canSendMessage(messageToSend messageToSend) bool {
	return messageToSend.sender.email != "" &&
		messageToSend.sender.phone != "" &&
		messageToSend.recipient.email != "" &&
		messageToSend.recipient.phone != ""
}

func main() {

	recipient := user{
		name:  "Jane",
		email: "coco@co.co",
		phone: "+123123123123",
	}

	messageToSendWrong := messageToSend{
		message: "Check this out",
		sender: user{
			name:  "John",
			email: "",
			phone: "",
		},
		recipient: recipient,
	}

	messageToSendCorrect := messageToSend{
		message: "Check this out",
		sender: user{
			name:  "John",
			email: "caca@ca.ca",
			phone: "+09876561565",
		},
		recipient: recipient,
	}

	checkMessage(messageToSendWrong)
	checkMessage(messageToSendCorrect)
}

func checkMessage(messageToSend messageToSend) {
	if canSendMessage(messageToSend) {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
