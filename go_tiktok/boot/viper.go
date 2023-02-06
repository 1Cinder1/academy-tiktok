package boot

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go_tiktok/app/global"
	"os"
)

const (
	configEnv  = "TIKTOK_CONFIG_PATH"
	configFile = "manifest/config/config.yaml"
)

func ViperSetup(path ...string) {
	var configPath string
	if len(path) != 0 {
		configPath = path[0]
	} else {
		flag.StringVar(&configPath, "c", "", "set config path")
		flag.Parse()

		if configPath == "" {
			if configPath = os.Getenv(configEnv); configPath != "" {

			} else {

				configPath = configFile
			}
		}
	}

	fmt.Printf("get config path: %s", configPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("get config file failed,err:%v", err))
	}
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("unmarshal config failed,err: %v", err))
	}
}
