package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:xxx@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil{
		return
	}
	return DB.DB().Ping()
}

func ModelAutoMigrate() {
	DB.AutoMigrate(&User{}, &Video{}, &Comment{})
}

func Close() error {
	err := DB.Close()
	return err
}