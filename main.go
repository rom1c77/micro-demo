package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"micro-demo/handler"
	pb "micro-demo/proto"
)

func main() {
	//go run main.go --registry=etcd
	//注意，这里我们手动通过 --registry=etcd 指定注册中心为 etcd，否则默认的注册中心是 mdns（在 Windows 系统中默认不可用）。
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("micro-demo"),
		micro.Version("latest"),
		// 注册consul中心
		micro.Registry(reg),
	)

	// Register handler
	if err := pb.RegisterMicroDemoHandler(srv.Server(), new(handler.MicroDemo)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
