package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"time"
)

func main() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "public", // 当存在多个 Namespace 时填写对应 Namespace ID，否则使用 public
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "data/nacos/log",
		CacheDir:            "data/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "mse-a49bd920-p.nacos-ans.mse.aliyuncs.com", // Nacos 服务的IP地址
			Port:   8848,
		},
	}

	var err error
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: serverConfig,
			ClientConfig:  &clientConfig})
	if err != nil {
		panic(err)
	}

	for {
		instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
			ServiceName: "gin-server-a.default.svc.cluster.local",
		})
		if err != nil {
			fmt.Printf("select one healthy instance err: %v\n", err)
		} else {
			fmt.Printf("selected instance info: %v\n", instance)
		}

		time.Sleep(time.Second * 5)
	}
}
