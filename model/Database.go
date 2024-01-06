package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	dsn := "host=localhost user=postgres password=12345 dbname=fiber-postgre port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err.Error())
	}

	
	fmt.Println("connected to postgre:",db)

	db.AutoMigrate(&Usertype{})
	db.AutoMigrate(&Users{})
	DB = db

}