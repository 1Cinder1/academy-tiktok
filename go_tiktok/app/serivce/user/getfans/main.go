package getfans

import (
	"go_tiktok/app/serivce/user/getfans/router"
	"go_tiktok/boot"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.RedisSetup()
	boot.MysqlSetup()
	boot.NacosSetUp("tiktok-getfans", 8081)
	router.InitRouter()

}
