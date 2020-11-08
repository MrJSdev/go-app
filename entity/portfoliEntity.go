package entity

// Hour struct to design SQL and respose
// Portfolio type
type Portfolio struct {
	ID          int    `json: ID`
	Company     string `json: Company`
	Developed   string `json: Developed`
	WorkOn      string `json: WorkOn`
	Skills      string `json: Skills`
	Description string `json: Description`
	Packages    string `json: Packages`
	Video       string `json: Video`
	Website     string `json: Website`
	AppUrl      string `json: AppUrl`
}
