package bot

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   string `gorm:"primaryKey"`
	Username string
}

type Tournament struct {
	ID uint `gorm:"primaryKey"`
}

func DatabaseConnect(databaseURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(databaseURL))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}
