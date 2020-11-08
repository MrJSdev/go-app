package entity

// Message type
type Message struct {
	ID     int    `json:"ID"`
	Text   string `json:"Text"`
	Sender string `json:"Sender"`
}
