package logic

type Message struct {
	// Message request type
	Action string `json:"action"`

	// Actual Message
	Message string `json:"message"`

	// Message target (room or user)
	Target string `json:"target"`

	// User sending the message
	Sender *User `json:"sender"`
}
