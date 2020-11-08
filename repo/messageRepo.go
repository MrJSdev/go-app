package repo

import (
	"MrHour/db"
	"MrHour/entity"
	"fmt"
	"log"

	// sql for connection
	_ "database/sql"
)

// FindMessages to list the messages
func FindMessages() ([]entity.Message, error) {
	DB, errDB := db.GetDB()
	if errDB != nil {
		fmt.Println("Error of connecting db")
	}
	var messages []entity.Message
	results, queryErr := DB.Query("SELECT * FROM messages")
	if queryErr != nil {
		fmt.Println("error while fethings hours list", queryErr)
	}
	for results.Next() {
		var (
			ID     int
			Text   string
			Sender string
		)
		err3 := results.Scan(&ID, &Text, &Sender)
		if err3 != nil {
			fmt.Println(err3.Error())
		} else {
			message := entity.Message{ID: ID, Text: Text, Sender: Sender}
			messages = append(messages, message)
		}
	}
	defer DB.Close()
	return messages, nil
}

// SendMessage to save message
func SendMessage(message entity.Message) error {
	DB, errDB := db.GetDB()
	if errDB != nil {
		fmt.Println("Error of connecting db")
	}
	fmt.Println(message, "message came from request")
	// h := entity.Hour{}
	stmt, err := DB.Prepare("INSERT INTO messages(Text, Sender) VALUES( ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	result, err := stmt.Exec(message.Text, message.Sender)
	fmt.Println(&result, "< resule")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
