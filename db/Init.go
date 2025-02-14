package db

import (
	"fmt"
	"ledger-system/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbConfig := config.LoadConfig()
	dbURL := dbConfig.GetDSN()

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connected successfully")

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
			fmt.Println("Database migrated successfully")
		case "seed":
			Seed(database)
			fmt.Println("Database seeded successfully")
		case "revert":
			Revert(database)
			fmt.Println("Database reverted successfully")
		}
	}
}
