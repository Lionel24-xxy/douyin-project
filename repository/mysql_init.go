package repository

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:xxy123456@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func ModelAutoMigrate() {
	DB.AutoMigrate(&User{}, &Video{}, &Comment{})
}

func Close() error {
	err := DB.Close()
	if err != nil {
		return errors.New("can't close current db")
	}
	return nil
}
