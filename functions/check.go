package functions

import (
	"fmt"
	"path/filepath"
	vars "pc-monitoring/config"
	"pc-monitoring/models/config"
	"pc-monitoring/database"
)

func CheckCompanies() {
    files := Filenames(vars.DataFolder, "*", "db")
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

func LoadCompanies() []config.Company {
    files := Filenames(vars.DataFolder, "*", "db")
    companies := make([]config.Company, 0)

    for _, file := range files {
        result := &database.Info{}

        db := database.Connect(file)
        query := db.First(&result)

        if query.Error == nil {
            companies = append(companies, config.Company{
                Name: result.Name,
                File: file,
            })
        } else {
            fmt.Printf("Cannot Get Company Information for %v\n", file)

            companies = append(companies, config.Company{
                Name: file,
                File: file,
            })
        }
    }

    return companies
}

func Filenames(folder, name, ext string) []string {
    matches, err := filepath.Glob(folder + "/" + name + "." + ext)

    if err != nil {
        return nil
    }

    return matches
}
