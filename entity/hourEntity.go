package entity

import "time"

// Hour struct to design SQL and respose
type Hour struct {
	ID        int       `json: ID`
	Date      time.Time `json: Date`
	Duration  string    `json: Duration`
	Note      string    `json: Note`
	CreatedAt time.Time `json: CreatedAt`
	UpdatedAt time.Time `json: UpdatedAt`
}
