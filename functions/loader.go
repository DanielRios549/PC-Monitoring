package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
    "strconv"

	// "pc-monitoring/models"
	"pc-monitoring/models"
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

func LoadAPConfig(configFile string) []*models.APResponse {
    var items []*models.APResponse

    loadConfig(configFile, "AP", func(device config.SnmpDevice, data []*config.Oid) {
        item := &models.APResponse{
            ID: device.ID,
            Hostname: "",
            Model: "",
            Version: "",
            Devices: "",
        }
    
        for _, option := range data {
            switch option.Name {
                case "hostname":
                    item.Hostname = option.Value
                case "ap_model":
                    item.Model = option.Value
                case "version":
                    item.Version = option.Value
                case "devices_count":
                    item.Devices = option.Value
            }
        }

        items = append(items, item)
    })

    return items
}

func LoadPrinterConfig(configFile string) []*models.PrinterResponse {
    var items []*models.PrinterResponse

    loadConfig(configFile, "printer", func(device config.SnmpDevice, data []*config.Oid) {
        item := &models.PrinterResponse{
            ID: device.ID,
            Hostname: "",
            Model: "",
            Toner_Percent: 0,
        }

        var toner_usage int
        var toner_max int
    
        for _, option := range data {

            switch option.Name {
                case "hostname":
                    item.Hostname = option.Value
                case "printer_model":
                    item.Model = option.Value
                case "toner_current":
                    valueNum, err := strconv.Atoi(option.Value)

                    if err != nil {
                        log.Fatalf("Failed to convert Toner Current: %v", err)
                    }

                    toner_usage = valueNum
                case "toner_max":
                    valueNum, err := strconv.Atoi(option.Value)

                    if err != nil {
                        log.Fatalf("Failed to convert Toner Max: %v", err)
                    }

                    toner_max = valueNum
            }
        }

        // Calculate Toner Percentage
        item.Toner_Percent = float32((toner_max - toner_usage) * 100)

        items = append(items, item)
    })

    return items
}

func loadConfig(configFile string, key string, callback func(device config.SnmpDevice, data []*config.Oid)) {
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

        var err error

        // By default it's a Printer
        roomItems := room.Printers
        isAP := key == "AP"

        if isAP {
            roomItems = room.APs
        }

		for _, item := range roomItems {
            version := item.Snmp.Version
			fmt.Printf("Item ID (V%d): %s\n", version, item.ID)

            var info []*config.Oid

            switch version {
                case 1:
                    info, err = snmp.V1(item.IP, isAP)
                case 2:
                    info, err = snmp.V2(item.IP, isAP)
                case 3:
                    info, err = snmp.V3(
                        item.IP,
                        item.Snmp.Context,
                        item.Snmp.User,
                        item.Snmp.Pass,
                        item.Snmp.Privpass,
                        isAP,
                    )
            }

            callback(item, info)

            if err != nil {
                log.Fatalf("Error Getting SNMP Info: %v", err)
            }
		}
	}
}
