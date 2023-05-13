package configs

import (
	"prakerja/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func AutoMigration() {
	DB.AutoMigrate(
		&models.Users{},
		&models.ExpenseCategories{},
		&models.IncomeCategories{},
		&models.Transactions{},
	)
}
