package database
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"auth-service/internal/config"
)

func connect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

}
