package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// idea db.Set("gorm:auto_preload", true)
// https://stackoverflow.com/questions/56832529/gorm-many-to-one-returns-empty
func ConnectToDb() {
  var err error
  
  dbUri := os.Getenv("DB_URI")
  DB, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect to the database")
  }
}
