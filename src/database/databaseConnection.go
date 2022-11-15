package database

import (
	"app/src/initialize"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var dsn = initialize.Env("DB_POSTGRES")
var DB *gorm.DB

func ConnectToDB() (err error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("DB connection error")
	}

	fmt.Println("DB connection success")

	DB = db
	return
}
