package model

type Manager struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `gorm:"varchar(20);not null"`
	Telephone      string `gorm:"varchar(20);not null;unique"`
	Password       string `gorm:"varchar(20);not null"`
	IsSuperManager bool   `gorm:"boolean;default false"`
}
