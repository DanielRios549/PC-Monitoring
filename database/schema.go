package database

import (
    "gorm.io/gorm"
)

type Company struct {
    ID            uint      `gorm:"primarykey"`
    Name          string    `gorm:"unique"`
    gorm.Model
}
