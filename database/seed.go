package database

var db = Connect()

func info() error {
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

func pc() error {
    pcs := []PC{
        {ID: 1, Name: "TI"},
    }

    for _, pc := range pcs {
        find := PC{
            Name: pc.Name,
        }

        result := db.Where(find).FirstOrCreate(&pc)

        if result.Error != nil {
            return result.Error
        }
    }

    return nil
}

func ap() error {
    aps := []AP{
        {ID: 1},
    }

    for _, ap := range aps {
        find := AP{
            ID: ap.ID,
        }

        result := db.Where(find).FirstOrCreate(&ap)

        if result.Error != nil {
            return result.Error
        }
    }

    return nil
}

func printer() error {
    printers := []AP{
        {ID: 1},
    }

    for _, printer := range printers {
        find := Printer{
            ID: printer.ID,
        }

        result := db.Where(find).FirstOrCreate(&printer)

        if result.Error != nil {
            return result.Error
        }
    }

    return nil
}

func Seed() error {
    tasks := []func() error{
		info,
		pc,
		ap,
        printer,
	}

	for _, task := range tasks {
        err := task()

		if err != nil {
            return err
		}
	}

    return nil
}
