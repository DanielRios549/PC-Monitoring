package database

func Seed() error {
    db := Connect()

    infos := []Info{
        {ID: 1, Name: "Main"},
    }

    for _, info := range infos {
        find := Info{
            Name: info.Name,
        }

        result := db.Where(find).FirstOrCreate(&info)

        if result.Error != nil {
            return result.Error
        }
    }

    return nil
}
