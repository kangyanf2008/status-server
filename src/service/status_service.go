package service

import (
	"status-server/protobuffer_def"
)

type StatusService interface {
	RegisterStatus(baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error
	QueryStatus(baseRequest *protobuffer_def.BaseRequest, baseResponse *protobuffer_def.BaseResponse) error
}
