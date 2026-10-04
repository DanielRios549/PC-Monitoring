package database

import (
	"pc-monitoring/config"
	"strings"

	"github.com/libtnb/sqlite"
	"gorm.io/gorm"
)

func Connect(filename ...string) *gorm.DB {
    file := config.DataFolder + "/main.db"
    
    if len(filename) >= 1 && strings.HasSuffix(filename[0], ".db") {
        file = filename[0]
    }

    db, err := gorm.Open(
        sqlite.Open(file),
        &gorm.Config{},
    )

    if err != nil {
        println("failed to connect database", err)
    }

    return db
}