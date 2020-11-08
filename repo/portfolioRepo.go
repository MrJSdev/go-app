package repo

import (
	"MrHour/db"
	"MrHour/entity"
	"fmt"
	"log"
)

// FindPortfolios to list the messages
func FindPortfolios() ([]entity.Portfolio, error) {
	portfolios := []entity.Portfolio{}
	// portfolio1 := entity.Portfolio{ID: 12, Company: "Emirates Transport", Skills: []string{
	// 	"VueJS",
	// 	"Element UI",
	// 	"Google Map",
	// 	"Javascript",
	// 	"SCSS",
	// 	"CSS",
	// 	"HTML"}, WorkOn: "Fleet Management App"}
	portfolio1 := entity.Portfolio{ID: 12, Company: "Emirates Transport", Skills: "Vue.js", WorkOn: "Fleet Management App"}
	portfolio2 := entity.Portfolio{ID: 13, Company: "DOT Transport", Skills: "Vue.js", WorkOn: "Fleet Management App"}
	portfolios = append(portfolios, portfolio1)
	portfolios = append(portfolios, portfolio2)
	return portfolios, nil
}

// SavePortfolio to save message
func SavePortfolio(portfolio entity.Portfolio) error {
	DB, errDB := db.GetDB()
	if errDB != nil {
		fmt.Println("Error of connecting db")
	}
	fmt.Println(portfolio, "portfolio came from request")
	stmt, err := DB.Prepare("INSERT INTO works( Company, WorkOn, Description, Developed, Skills, Packages, Video, Website, AppUrl) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(portfolio.Company, portfolio.WorkOn, portfolio.Description, portfolio.Developed, portfolio.Skills, portfolio.Packages, portfolio.Video, portfolio.Website, portfolio.AppUrl)
	fmt.Println(&result, "< resule")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
