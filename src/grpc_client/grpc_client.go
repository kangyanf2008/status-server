package grpc_client

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"status-server/protobuffer_def"
	"time"
)

func Start()  {
	r := zookeeper.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2181"}
		op.Context = context.Background()
		op.Timeout = time.Second * 5
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// create GRPC service
	service := grpc.NewService(
		service.Name("test.client2"),
		service.Registry(r),
		service.Context(ctx),

	)

	service.Client().Init(client.Retries(3),client.PoolSize(200), client.PoolTTL(time.Second*20), client.RequestTimeout(time.Second*5))

	test :=  protobuffer_def.NewStatusServerService("status-service", service.Client())

	for r:=0; r<20; r++ {
		go func() {
			i := 0
			for {
				_, err := test.BaseInterface(context.Background(), &protobuffer_def.BaseRequest{RequestId:"1111",C:protobuffer_def.CMD_REGISTER_STATUS})
				if err != nil {
					fmt.Println(err)
				} else {
					i ++
				}
				if i % 10000 == 0 {
					fmt.Println(i, time.Now().Unix())
				}
			}
		}()
	}


}