package model

type Student struct {
	ID    uint   `gorm:"primarykey"`
	Sid   string `gorm:"varchar(20);not null;unique"`
	Name  string `gorm:"varchar(20);not null"`
	Sex   string `gorm:"varchar(5);not null"`
	Age   int    `gorm:"integer(5);not null"`
	Major string `gorm:"varchar(36);not null"`
	Class string `gorm:"varchar(20);not null"`
}
