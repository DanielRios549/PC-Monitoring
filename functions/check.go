package functions

import (
	"fmt"
	"path/filepath"
	"pc-monitoring/database"
)

func CheckCompanies() {
    files := filenames("data", "*", "db")
    count := len(files)

    if count < 1 {
        err := database.Migrate()

        if err != nil {
            panic("Error to Migrate First Database")
        }

        err = database.Seed()

        if err != nil {
            panic("Error to Seed First Company")
        }

        fmt.Println("Created first Database file")
    }
}

func filenames(folder, name, ext string) []string {
    matches, err := filepath.Glob(folder + "/" + name + "." + ext)

    if err != nil {
        return nil
    }

    return matches
}
