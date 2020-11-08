package repo

import (
	"MrHour/db"
	"MrHour/entity"
	"log"

	// sql for connection
	_ "database/sql"
	"fmt"
	"time"
)

// type repo struct{}

// // NewSQLRepositor constructor for message query
// func NewSQLRepositor() HourRepo {
// 	return &repo{}
// }

// FindAll get all hours
func FindAll() ([]entity.Hour, error) {
	DB, errDB := db.GetDB()
	if errDB != nil {
		fmt.Println("Error of connecting db")
	}
	var hours []entity.Hour
	results, queryErr := DB.Query("SELECT * FROM works")
	if queryErr != nil {
		fmt.Println("error while fethings hours list", queryErr)
	}
	for results.Next() {
		var (
			id         int
			note       string
			created_at time.Time
			updated_at time.Time
			date       time.Time
			duration   string
		)
		err3 := results.Scan(&id, &created_at, &updated_at, &date, &duration, &note)
		if err3 != nil {
			fmt.Println(err3.Error())
		} else {
			hour := entity.Hour{ID: id, Note: note, UpdatedAt: updated_at, CreatedAt: created_at, Duration: duration, Date: date}
			hours = append(hours, hour)
		}
	}
	defer DB.Close()
	return hours, nil
}

// Save each hour
func Save(data entity.Hour) (entity.Hour, error) {
	DB, errDB := db.GetDB()
	if errDB != nil {
		fmt.Println("Error of connecting db")
	}
	fmt.Println(data, "data came from request")
	// h := entity.Hour{}
	stmt, err := DB.Prepare("INSERT INTO works(note, updated_at, created_at, duration, date) VALUES( ?, ?, ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	result, err := stmt.Exec(data.Note, data.UpdatedAt, data.CreatedAt, data.Duration, data.Date)
	fmt.Println(&result, "< resule")
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
