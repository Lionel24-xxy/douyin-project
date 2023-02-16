package main

import (
	"TikTok_Project/repository"
	"TikTok_Project/router"
	"log"
)

func main() {
	err := repository.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := repository.Close()
		if err != nil {
			log.Println("can't close current db！")
		}
	}()
	repository.ModelAutoMigrate()

	r := router.InitRouter()
	r.Run()
}
