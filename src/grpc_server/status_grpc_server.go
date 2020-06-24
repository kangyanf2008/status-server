package grpc_server

import (
	"context"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
	"status-server/config"
	"status-server/protobuffer_def"
	"status-server/service/impl"
	"sync"
	"time"
)

var (
	statusServers service.Service
	statusOnce    = &sync.Once{}
	statusServer  server.Server
)

//初始化grpc service服务
func StartStatusService(config *config.Config) {
	statusOnce.Do(func() {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// create GRPC service
		service := grpc.NewService(
			service.Address(config.Base.GRPCAddr),
			service.Name(config.Base.ServiceName),
			service.Registry(config.RegisterCenter.GetRegisterCenter()),
			service.RegisterTTL(time.Second*30),
			service.RegisterInterval(time.Second*20),
			service.Context(ctx),
		)

		service.Init()

		statusServer = service.Server()
		// register test handler
		protobuffer_def.RegisterStatusServerHandler(service.Server(), &statusGrpcServiceImpl{})

		//启动服务
		if err := service.Run(); err != nil {
			panic(err)
		}
	})
}

func StopStatusService()  {
	statusServer.Stop()
}

type statusGrpcServiceImpl struct{}

func (s *statusGrpcServiceImpl) BaseInterface(context context.Context, baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error {
	baseResponse.C = baseRequest.GetC()
	baseResponse.RequestId = baseRequest.GetRequestId()
	baseResponse.Code = protobuffer_def.ReturnCode_SUCCESS

	switch baseRequest.GetC() {
	case protobuffer_def.CMD_REGISTER_STATUS:   //注册状态
		return impl.NewStatusService().RegisterStatus(baseRequest, baseResponse)
	case protobuffer_def.CMD_QUERY_STATUS:      //查询状态
		return impl.NewStatusService().QueryStatus(baseRequest, baseResponse)
	default:
		baseResponse.Desc = "unkown cmd"
		baseResponse.Code = protobuffer_def.ReturnCode_UNKOWN_CMD; //示知的指令
	}
	return nil
}

