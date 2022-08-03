package main

import (
	"flag"
	"strings"
	"time"

	user_agent "github.com/hwholiday/learning_tools/micro_agent/proto/user"

	"github.com/hwholiday/learning_tools/micro_agent/base"
	"github.com/hwholiday/learning_tools/micro_agent/base/config"
	"github.com/hwholiday/learning_tools/micro_agent/base/tool"
	"github.com/hwholiday/learning_tools/micro_agent/handler"
	"github.com/hwholiday/learning_tools/micro_agent/model"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var conf = flag.String("conf", "/home/ghost/go/src/micro_agent/conf", "conf path")

func main() {
	base.Init(*conf)
	registry := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Timeout = time.Second * 5
		options.Addrs = strings.Split(config.GetServerConfig().GetEtcdAddr(), ",")
	})
	service := micro.NewService(
		micro.Name(config.GetServerConfig().GetServerName()),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
		}),
	)
	// 注册服务
	tool.GetLogger().Info("start service " + config.GetServerConfig().GetServerName() + " success")
	_ = user_agent.RegisterUserHandler(service.Server(), handler.GetService())
	// 启动服务
	if err := service.Run(); err != nil {
		panic(err)
	}
}
