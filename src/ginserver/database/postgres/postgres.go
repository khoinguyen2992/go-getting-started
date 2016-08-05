package postgres

import (
	"fmt"
	"ginserver/env"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	if DB == nil {
		connectionStr := getConnectionString()
		db, err := gorm.Open("postgres", connectionStr)
		if err != nil {
			panic(err.Error())
		}

		db.DB()
		db.DB().Ping()
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		db.SingularTable(false)
		db.LogMode(true)
		DB = db
	}
}

func getConnectionString() string {
	connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
	return fmt.Sprintf(connectionStringTemplate,
		env.Get("DB_HOST"),
		env.Get("DB_PORT"),
		env.Get("DB_SSL_MODE"),
		env.Get("DB_USERNAME"),
		env.Get("DB_PASSWORD"),
		env.Get("DB_NAME"))
}
