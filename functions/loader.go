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

func LoadConfig(configFile string) [][]*config.Oid {
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

    var printers [][]*config.Oid
	
	for _, room := range floor.Rooms {
		fmt.Printf("Room Name: %s\n", room.Name)

        var err error

		for _, printer := range room.Printers {
            version := printer.Snmp.Version
			fmt.Printf("Printer ID (V%d): %s\n", version, printer.ID)

            var info []*config.Oid

            switch version {
                case 1:
                    info, err = snmp.V1(printer.IP)
                case 2:
                    info, err = snmp.V2(printer.IP)
                case 3:
                    info, err = snmp.V3(
                        printer.IP,
                        printer.Snmp.Context,
                        printer.Snmp.User,
                        printer.Snmp.Pass,
                        printer.Snmp.Privpass,
                    )
            }

            printers = append(printers, info)

            if err != nil {
                log.Fatalf("Error Getting SNMP Info: %v", err)
            }
		}
	}

    return printers
}
