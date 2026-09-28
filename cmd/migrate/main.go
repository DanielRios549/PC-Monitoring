package main

import (
    "pc-monitoring/database"
)

func main() {
    db := database.Connect()

    err := db.AutoMigrate(
        &database.Company{},
    )

    if err != nil {
        println("Error to Migrate: ", err)
    }
}
