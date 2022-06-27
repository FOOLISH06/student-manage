package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"student-manage/model"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		db = initDB()
	}
	return db
}

func initDB() *gorm.DB {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("fail to connect to database, err:v", err.Error())
	}

	if err := db.AutoMigrate(&model.Student{}); err != nil {
		log.Fatalln("fail to migrate table student, err: ", err.Error())
	}
	if err := db.AutoMigrate(&model.Manager{}); err != nil {
		log.Fatalln("fail to migrate table manager, err: ", err.Error())
	}

	return db
}
