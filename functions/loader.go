package functions

import (
	"log"
	"os"
	"pc-monitoring/models/plan"

	"github.com/joho/godotenv"
)

var Plan plan.FloorPlan

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
