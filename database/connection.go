package database

import (
	"Bookstore/objects"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	dsn := "postgres://bbaizhanov:YSd5aiBNjO4R@ep-floral-bush-657657.us-east-2.aws.neon.tech/neondb"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&objects.Book{})

	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		sleep := time.Duration(1)
		for dbase == nil {
			sleep *= 2
			fmt.Printf("Database is unavailable. Wait %d seconds", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
