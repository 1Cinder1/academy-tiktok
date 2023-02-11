package main

import (
	"go_tiktok/app/serivce/user/getfriends/router"
	"go_tiktok/boot"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.RedisSetup()
	boot.MysqlSetup()
	boot.NacosSetUp("tiktok-getfriends", 8081)
	router.InitRouter()

}
