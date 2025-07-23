package initializer

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
	var err error
	dsn := "host=ep-fragrant-bonus-adrua1lt-pooler.c-2.us-east-1.aws.neon.tech user=neondb_owner password=npg_tC74BTSsJhrZ dbname=neondb port=5432 sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return err
}
