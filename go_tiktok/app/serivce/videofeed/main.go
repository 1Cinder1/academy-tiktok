package videofeed

import (
	"go_tiktok/app/serivce/videofeed/router"
	"go_tiktok/boot"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.RedisSetup()
	boot.MysqlSetup()
	router.InitRouter()

}
