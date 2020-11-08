package db

import (
	"database/sql"
	"fmt"

	// mysql driver import
	_ "github.com/go-sql-driver/mysql"
)

// GetDB the database
func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "mrjs"
	dbUser := "root"
	dbPassword := "root"
	dbHost := "localhost"
	dbPort := ":3306"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+dbHost+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		fmt.Println("having problem to connnect mysql db", err)
	}
	return
}
