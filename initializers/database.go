package initializers

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

/*func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_ULR")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed connect to database")
	}
}*/
