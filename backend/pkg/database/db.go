package database

import (
	"log"
	"os"

	"niya-voyage/backend/pkg/models"

	//"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func ConnectDatabase() {
	//var database *gorm.DB
	var err error

	// db_hostname := os.Getenv("POSTGRES_HOST")
	// db_name := os.Getenv("POSTGRES_DB")
	// db_user := os.Getenv("POSTGRES_USER")
	// db_pass := os.Getenv("POSTGRES_PASSWORD")
	// db_port := os.Getenv("POSTGRES_PORT")

	// dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)

	// for i := 1; i <= 3; i++ {
	// 	database, err = gorm.Open(postgres.Open(dbURl), &gorm.Config{})
	// 	if err == nil {
	// 		break
	// 	} else {
	// 		log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
	// 		time.Sleep(3 * time.Second)
	// 	}
	// }
	db, err := gorm.Open(sqlite.Open("niya-voyage.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	//DB = database
	Database = DbInstance{
		DB: db,
	}
}
