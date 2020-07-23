package models

type Receiver struct {
	ID         int    `json:"id"`
	Action     string `json:"action"`
	DocumentID string `json:"document_id"`
}

type Receivers []Receiver
type ReceiverSuccess struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *Receiver `json:"data"`
}
