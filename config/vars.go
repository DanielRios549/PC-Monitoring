package config

import "os"

func getenv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
        return fallback
    }

    return value
}

var HOST = getenv("HOST", "localhost")
var PORT = getenv("PORT", "9003")
var DataFolder = getenv("DATA_FOLDER", "data")
