package main

import (
	"TikTok_Project/repository"
	"TikTok_Project/router"
	"TikTok_Project/utils"
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

	if err := utils.SensitiveWordInit(); err != nil {
		log.Printf("敏感词初始化失败")
		panic(err)
	}

	r := router.InitRouter()
	r.Run()
}
