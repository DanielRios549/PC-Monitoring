package database

import (
    "gorm.io/gorm"
    "pc-monitoring/models/config"
)

type Info struct {
    ID            uint      `gorm:"primarykey"`
    Name          string    `gorm:"unique"`
    gorm.Model
}

type PC struct {
    ID            uint      `gorm:"primarykey"`
    config.PC
    gorm.Model
}

type AP struct {
    ID            uint      `gorm:"primarykey"`
    config.SnmpDevice
    gorm.Model
}

type Printer struct {
    ID            uint      `gorm:"primarykey"`
    config.SnmpDevice
    gorm.Model
}

func Migrate() error {
    db := Connect()

    err := db.AutoMigrate(
        &Info{},
        &PC{},
        &AP{},
        &Printer{},
    )

    if err != nil {
        return err
    }

    return nil
}

