package getuser

import (
	"go_tiktok/app/serivce/user/getuser/router"
	"go_tiktok/boot"
)

func main() {
	boot.ViperSetup()
	boot.LoggerSetup()
	boot.RedisSetup()
	boot.MysqlSetup()
	router.InitRouter()

}
