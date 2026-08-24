package functions

import (
	"log"
	"os"
	"encoding/json"
	"fmt"
	"pc-monitoring/models/plan"
	"pc-monitoring/snmp"

	"github.com/joho/godotenv"
)

var Plan plan.Floor

func LoadEnv() {
	env := os.Getenv("ENV")

    if env == "" {
        env = "development"
    }

    err := godotenv.Load(".env." + env)

	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer func() {
		err := file.Close()

		if err != nil {
			log.Fatalf("Cannot Close Connection: %v", err)
		}
	}()

	var floor plan.Floor

	err = json.NewDecoder(file).Decode(&floor)

	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	fmt.Println("Successfully loaded config:")
	fmt.Printf("Floor Name: %s\n", floor.Name)
	
	for _, room := range floor.Rooms {
		fmt.Printf("Room Name: %s\n", room.Name)

		for _, printer := range room.Printers {
			fmt.Printf("Printer ID: %s\n", printer.ID)

			// TODO: Get version dynamic
			snmp.V3(
				printer.IP,
				printer.Snpm.Context,
				printer.Snpm.User,
				printer.Snpm.Pass,
				printer.Snpm.Privpass,
			)
		}
	}
}
