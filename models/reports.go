package models

type CategoryReport struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Total    int    `json:"total"`
}

type Report struct {
	TotalExpenses  int         `json:"totalExpenses"`
	TotalIncome    int         `json:"totalIncome"`
	Difference     int         `json:"difference"`
	CategoryReport interface{} `json:"categoryReport"`
}
