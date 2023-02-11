package boot

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"go.uber.org/zap"
	"go_tiktok/app/global"

	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NacosSetUp(serviceName string, servicePort uint64) {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{ // nacos服务的ip和端口及请求方式
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		global.Logger.Fatal("register service failed.", zap.Error(err))
	}

	//注册实例：RegisterInstance
	flag, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",                     //注册服务的ip
		Port:        servicePort,                     //注册服务的端口
		ServiceName: serviceName,                     //注册服务名
		Weight:      10,                              //权重
		Enable:      true,                            //是否可用
		Healthy:     true,                            //健康状态
		Ephemeral:   true,                            //零时节点（服务下线之后，nacos 上注册信息会删除）
		Metadata:    map[string]string{"name": "go"}, //元数据
		ClusterName: "DEFAULT",                       // 默认值DEFAULT  集群名称
		GroupName:   "DEFAULT_GROUP",                 // 默认值DEFAULT_GROUP 组名称
	})
	if flag {
		global.Logger.Info("register service successfully!")
	} else {
		global.Logger.Fatal("register service failed.", zap.Error(err))
	}
}
