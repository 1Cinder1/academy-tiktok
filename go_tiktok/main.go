package main

import (
	"go_tiktok/boot"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.RedisSetup()
	boot.MysqlSetup()

}
