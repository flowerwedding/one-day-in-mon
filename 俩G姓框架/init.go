package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/dome7?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	if !db.HasTable(&Mon{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Mon{}).Error; err != nil {
			panic(err)
		}
	}
}

type Mon struct {
	Id        int    `gorm:"type:int(255);not null;"`
	Username  string `gorm:"type:varchar(256);not null;"`
	Password  string `gorm:"type:varchar(256);not null;"`
	CreatedAt time.Time
}