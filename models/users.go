package models

type Users struct {
	Id       int    `json:"id"`
	Role     string `json:"role"`
	Balance  int    `json:"balance"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
