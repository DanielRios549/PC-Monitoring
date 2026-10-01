package main

import (
    "pc-monitoring/database"
)

func main() {
    db := database.Connect()

    err := db.AutoMigrate(
        &database.Company{},
        &database.PC{},
        &database.AP{},
        &database.Printer{},
    )

    if err != nil {
        println("Error to Migrate: ", err)
    }
}
