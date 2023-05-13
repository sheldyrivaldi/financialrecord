package models

type Transactions struct {
	Id       int    `json:"id"`
	User_ID  int    `json:"user_id"`
	Date     string `json:"date"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Amount   int    `json:"amount"`
	Note     string `json:"note"`
}
