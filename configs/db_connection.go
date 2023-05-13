package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	dsn := "4g4l823idwofed3o60db:pscale_pw_3auHMq5Beinpg72rQiXDdsrXrtPfRKj5egFXbnWhy5g@tcp(aws.connect.psdb.cloud)/finrecords?tls=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")
	}

	AutoMigration()
}
