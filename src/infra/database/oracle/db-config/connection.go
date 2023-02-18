package dbconfig

import (
	"fmt"
	"os"

	"github.com/cengsin/oracle"
	"gorm.io/gorm"
)

var Oracle *gorm.DB

func Connection() {
	dns := fmt.Sprintf("%s:%s@//%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(oracle.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	Oracle = db
}
