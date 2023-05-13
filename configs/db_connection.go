package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	dsn := "root:VDXdZxHQu7G9on1DGjMk@tcp(containers-us-west-168.railway.app:7498)/railway?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")
	}

	AutoMigration()
}
