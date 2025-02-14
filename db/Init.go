package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(args []string) *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/ledger_system?sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully")

	// Add run commands
	RunCommands(db, os.Args[1:])

	return db
}

func RunCommands(database *gorm.DB, args []string) {
	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}

	// Test the database connection
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	for _, arg := range args {
		switch arg {
		case "migrate":
			Migrate(database)
		case "seed":
			Seed(database)
		}
	}
}
