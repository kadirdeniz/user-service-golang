package tools

import (
	"fmt"
	"user-service-golang/pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetPostgresql() *gorm.DB {
	mutex.Lock()
	defer mutex.Unlock()

	if db == nil {
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
			pkg.Configs.Database.Host,
			pkg.Configs.Database.User,
			pkg.Configs.Database.Password,
			pkg.Configs.Database.DBName,
			pkg.Configs.Database.Port,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		return db
	}
	return db
}
