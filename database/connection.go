package database

import (
	"pc-monitoring/config"

	"github.com/libtnb/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

    db, err := gorm.Open(
        sqlite.Open(config.DataFolder + "/main.db"),
        &gorm.Config{},
    )

    if err != nil {
        println("failed to connect database", err)
    }

    return db
}