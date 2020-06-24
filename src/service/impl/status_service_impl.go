package impl

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"status-server/protobuffer_def"
	lredis "status-server/redis"
	serivce "status-server/service"
	"sync"
)

var (
	ssi *statusServiceImpl
	ssiOnce  = &sync.Once{}
)

type statusServiceImpl struct {}

func init()  {
	fmt.Println("init status")
}

func NewStatusService() serivce.StatusService {
	ssiOnce.Do(func() {
		ssi = &statusServiceImpl{}
	})
	return ssi
}

func  (s *statusServiceImpl)  preDeal(baseRequest *protobuffer_def.BaseRequest,
	baseResponse *protobuffer_def.BaseResponse, request proto.Message)  bool {
	//解析body
	if baseRequest.GetBody() == nil {
		baseResponse.Code = protobuffer_def.ReturnCode_BODY_IS_NULL
		baseResponse.Desc = "body is null"
		return false
	}
	//反序列化

	err := ptypes.UnmarshalAny(baseRequest.GetBody(), request)
	if err != nil {
		baseResponse.Code = protobuffer_def.ReturnCode_DESERIALIZATION_ERROR
		baseResponse.Desc = "body deserialization error"
		return false
	}
	return true
}

//注册状态
func (s *statusServiceImpl) RegisterStatus(baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error {
	baseResponse.Code = protobuffer_def.ReturnCode_SUCCESS

	//解析请求参数
	request := &protobuffer_def.RegisterStatusRequest{}
	if !s.preDeal(baseRequest, baseResponse, request) {
		return nil
	}

	//注册状态
	//err := lredis.RegisterStatus(request)
	err := lredis.RegisterStatus2(request)
	if err != nil {
		baseResponse.Code = protobuffer_def.ReturnCode_UNKOWN_ERROR
		baseResponse.Desc = "save status err"
		return err
	}
	return nil
}

//查询状态
func (s *statusServiceImpl) QueryStatus(baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error {
	baseResponse.Code = protobuffer_def.ReturnCode_SUCCESS

	//解析请求参数
	request := &protobuffer_def.QueryStatusRequest{}
	if !s.preDeal(baseRequest, baseResponse, request) {
		return nil
	}

	//缓存中查询用户状态信息
	response, err := lredis.QueryStatus(request)
	if err != nil {
		baseResponse.Code = protobuffer_def.ReturnCode_UNKOWN_ERROR
		baseResponse.Desc = "save status error"
		return err
	}

	//序列化状态信息
	if response != nil {
		body, err2 :=ptypes.MarshalAny(response)
		if err2 != nil {
			baseResponse.Code = protobuffer_def.ReturnCode_SERIALIZATION_ERROR
			baseResponse.Desc = "MarshalAny error"
		}
		baseResponse.Body = body
		return err2
	}
	return nil
}

