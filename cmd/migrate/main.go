package main

import (
    "pc-monitoring/database"
)

func main() {
    err := database.Migrate()

    if err != nil {
        println("Error to Migrate: ", err)
    }
}
