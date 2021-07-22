package message

type Message struct {
	Type int    `json:"type"`
	ID string `json:"id"`
	Body string `json:"body"`
}
