package main

import (
    "pc-monitoring/database"
)

func main() {
    err := database.Seed()

    if err != nil {
        println("Error to Seed: ", err)
    }
}
