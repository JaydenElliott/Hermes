package logic

import (
	"encoding/json"
	"log"
)

type Message struct {
	// Message request type
	Action string `json:"action"`

	// Actual Message
	Message string `json:"message"`

	// Message target (room or user)
	Target *Channel `json:"target"`

	// User sending the message
	Sender *User `json:"sender"`
}

func (message *Message) marshal() []byte {
	_json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return _json
}
