package models

type ExpenseCategories struct {
	Id       int    `json:"id"`
	User_ID  int    `json:"user_id"`
	Category string `json:"category"`
}
