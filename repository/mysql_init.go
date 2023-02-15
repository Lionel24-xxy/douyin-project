package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	//xxy
	//dsn := "root:xxy123456@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	//Leotao
	dsn := "root:123456@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func ModelAutoMigrate() {
	DB.AutoMigrate(&User{}, &Video{}, &Comment{})
}

func Close() {
	DB.Close()
}
