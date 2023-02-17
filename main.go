package main

import (
	"TikTok_Project/router"
	"TikTok_Project/repository"
)



func main() {
	err := repository.InitMySQL()
	if err != nil{
		panic(err)
	}
	defer repository.Close()
	repository.ModelAutoMigrate()

	r := router.InitRouter()
	r.Run()
}