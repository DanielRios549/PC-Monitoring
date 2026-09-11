package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	// "pc-monitoring/models"
	"pc-monitoring/models/config"
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

func LoadConfig(configFile string, key string) [][]*config.Oid {
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

    var items [][]*config.Oid

	for _, room := range floor.Rooms {
		fmt.Printf("Room Name: %s\n", room.Name)

        var err error

        // By default it's a Printer
        keyItems := room.Printers

        if key == "AP" {
            keyItems = room.APs
        }

		for _, item := range keyItems {
            version := item.Snmp.Version
			fmt.Printf("Item ID (V%d): %s\n", version, item.ID)

            var info []*config.Oid

            switch version {
                case 1:
                    info, err = snmp.V1(item.IP)
                case 2:
                    info, err = snmp.V2(item.IP)
                case 3:
                    info, err = snmp.V3(
                        item.IP,
                        item.Snmp.Context,
                        item.Snmp.User,
                        item.Snmp.Pass,
                        item.Snmp.Privpass,
                    )
            }

            items = append(items, info)

            if err != nil {
                log.Fatalf("Error Getting SNMP Info: %v", err)
            }
		}
	}

    return items
}
