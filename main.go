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

	if err := repository.InitRedisClient(); err != nil {
		panic(err)
	}

	r := router.InitRouter()
	r.Run()
}
