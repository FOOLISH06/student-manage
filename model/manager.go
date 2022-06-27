package model

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null"`
	Telephone string `gorm:"varchar(20);not null;unique"`
	Password  string `gorm:"varchar(20);not null"`
}
