package redis

import (
	"github.com/go-redis/redis"
	"status-server/constants"
	"status-server/protobuffer_def"
	"status-server/utils"
	"strconv"
	"strings"
	"time"
)


type statusInfoStruct struct {
	deviceType               byte
	statusInfo				 string
	lastTime                 int64
	interval				 int32
}


func RegisterStatus(request *protobuffer_def.RegisterStatusRequest) error {
	pipeline := Pipeline()
	defer pipeline.Close()

	identify := request.GetIdentity()                    //唯一标识
	interval := request.GetNextHeartbeatInterval()       //下次心跳间隔
	deviceTypeBytes := request.GetDeviceType()           //设备类型
	registerInfo := request.GetRegisterInfo()            //注册信息

	var deviceType byte
	if len(deviceTypeBytes) == 0 {
		deviceType = 0
	} else {
		deviceType = deviceTypeBytes[0]
	}

	hmapkey := constants.GetIdentityKey(identify)

	lastTimekey := constants.GetLastTimeKey(deviceType)                    //最后一次次心跳时间key
	heartbeatInterval := constants.GetNextHeartbeatIntervalKey(deviceType) //心跳时间间隔
	registerInfoKey := constants.GetRegisterInfoKey(deviceType)            //心跳时间间隔

	value := make(map[string]interface{},3)
	value[lastTimekey] = utils.TimeStamp()
	value[heartbeatInterval] = interval
	value[registerInfoKey] = registerInfo

	//上次心跳时间
	pipeline.HMSet(hmapkey, value)

	//状态标记, DELAY_SECOND_TIME延迟时间
	pipeline.Expire(hmapkey, time.Second*time.Duration(interval) + constants.DELAY_SECOND_TIME)

	//执行redis查询
	_, e := pipeline.Exec()
	return e
}

func RegisterStatus2(request *protobuffer_def.RegisterStatusRequest) error {

	identify := request.GetIdentity()                    //唯一标识
	interval := request.GetNextHeartbeatInterval()       //下次心跳间隔,单位秒
	deviceTypeBytes := request.GetDeviceType()           //设备类型
	registerInfo := request.GetRegisterInfo()            //注册信息

	var deviceType byte
	if len(deviceTypeBytes) == 0 {
		deviceType = 0
	} else {
		deviceType = deviceTypeBytes[0]
	}

	hMapkey := constants.GetIdentityKey(identify)

	lastTimekey := constants.GetLastTimeKey(deviceType)                       //最后一次次心跳时间key
	heartbeatIntervalKey := constants.GetNextHeartbeatIntervalKey(deviceType) //心跳时间间隔
	registerInfoKey := constants.GetRegisterInfoKey(deviceType)               //心跳时间间隔
	//DELAY_SECOND_TIME延迟时间
	expire := interval + constants.DELAY_SECOND_TIME

	_,e := StatusLua.EvalSha([]string{hMapkey}, lastTimekey, utils.TimeStamp(), heartbeatIntervalKey, interval, registerInfoKey, registerInfo, expire)
	return  e
}


func QueryStatus(request *protobuffer_def.QueryStatusRequest) (*protobuffer_def.QueryStatusResponse, error) {

	identify := request.GetIdentity()                    //唯一标识
	hMapkey := constants.GetIdentityKey(identify)

	var r *redis.StringStringMapCmd
	if Model == constants.RedisModel {
		r = ClusterClient.HGetAll(hMapkey)
	} else {
		r = Client.HGetAll(hMapkey)
	}
	if r != nil {
		if r.Err() != nil {
			return  nil, r.Err()
		}
		//安设备类型存放到map中
		statusInfo := make(map[byte]*statusInfoStruct)
		for key, value := range r.Val() {
			keys := strings.Split(key,constants.Colon)
			deviceType := keys[0]
			subKey := keys[1]
			typeInt,_ := strconv.Atoi(deviceType)
			if v, ok := statusInfo[byte(typeInt)]; ok {
				assembleStatusInfo(subKey, value, v)
			} else {
				s := &statusInfoStruct{deviceType:byte(typeInt),}
				statusInfo[byte(typeInt)] = s
				assembleStatusInfo(subKey, value, s)
			}
		}
		//过滤到未过期的状态信息
		var resultStatusInfo []*protobuffer_def.QueryStatusResponse_StatusInfo
		currentTime := utils.TimeStamp()
		for _,v := range statusInfo {
			//判断是否已经过期
			if (v.lastTime + int64(v.interval) * 1000) >= currentTime {
				resultStatusInfo = append(resultStatusInfo, &protobuffer_def.QueryStatusResponse_StatusInfo{DeviceType:[]byte{v.deviceType}, RegisterInfo:v.statusInfo})
			} else {
				//TODO 删除无效的缓存数据，需要使用分布式锁
			}
		}
		//如果不为空，则返回用户状态信息
		if len(resultStatusInfo) > 0 {
			return  &protobuffer_def.QueryStatusResponse{Status:resultStatusInfo,}, nil
		}
		return  nil, nil
	}

	return nil, nil
}

//组装信息
func assembleStatusInfo(subKey, value string, s *statusInfoStruct) {
	switch subKey {
	case constants.LastTimeKeySubKey:
		s.lastTime, _ = strconv.ParseInt(value,10, 64)
	case constants.NextHeartbeatIntervalSubKey:
		int64Interval, _:= strconv.ParseInt(value,10, 32)
		s.interval = int32(int64Interval)
	case constants.RegisterInfoSubKey:
		s.statusInfo = value
	}
}

/*//保存用户状态
func UpdateUserStatus(request *protobuffer.StatusServerUpdateStatusRequest) error {
	pipeline := Pipeline()
	defer pipeline.Close()

	strUserid := fmt.Sprint(request.GetUserid())
	hsetkey := constants.GetUserStatusHSetKey(strUserid)
	//上次心跳时间key
	lastTimekey := constants.GetUserLastTimeKey()

	//上次心跳时间
	pipeline.HSet(hsetkey, lastTimekey, utils.TimeStamp())

	//状态key
	serverIpKey := constants.GetUserAccessIpKey(strUserid, request.GetTType())
	clientIpKey := constants.GetUserClientIpKey(strUserid, request.GetTType())
	//deviceKey := constants.GetUserDeviceKey(strUserid, request.GetTType())

	//状态标记
	pipeline.Set(serverIpKey, request.GetServerIp(), time.Second*time.Duration(request.GetNextRequest())+constants.DELAY_SECOND_TIME)
	//客户端IP
	pipeline.Set(clientIpKey, request.GetClientIp(), time.Second*time.Duration(request.GetNextRequest())+constants.DELAY_SECOND_TIME)
	//在线设备唯一ID
	//pipeline.Set(deviceKey, request.GetDeviceId(), time.Second*time.Duration(request.GetNextRequest())+constants.DELAY_SECOND_TIME)

	//执行redis查询
	_, e := pipeline.Exec()
	if e != nil {
		return e
	}
	return e
}

//清除用户在线状态
func DelUserStatus(request *protobuffer.StatusServerUpdateStatusRequest) error {
	pipeline := Pipeline()
	defer pipeline.Close()

	strUserId := fmt.Sprint(request.GetUserid())
	//状态key
	serverIpKey := constants.GetUserAccessIpKey(strUserId, request.GetTType())
	clientIpKey := constants.GetUserClientIpKey(strUserId, request.GetTType())
	//deviceKey := constants.GetUserDeviceKey(strUserId, request.GetTType())

	//服务器接入层Ip
	pipeline.Del(serverIpKey)
	//客户端IP
	pipeline.Del(clientIpKey)
	//在线设备唯一ID
	//pipeline.Del(deviceKey)

	//执行redis查询
	_, e := pipeline.Exec()
	if e != nil {
		return e
	}
	return e
}

//注册用户在线状态
func RegisterUserStatus(request *protobuffer.StatusServerRegisterStatusRequest) error {
	pipeline := Pipeline()
	defer pipeline.Close()

	strUserId := fmt.Sprint(request.GetUserid())
	//用户状态hset key
	hsetkey := constants.GetUserStatusHSetKey(strUserId)
	//用户SDK 类型
	sdkType := constants.GetSdkTypeKey()
	//上次心跳时间key
	lastTimekey := constants.GetUserLastTimeKey()
	deviceKey := constants.GetUserDeviceKey(protobuffer.TerminalType_MOBILE)

	//上次心跳时间和sdk版本号
	pipeline.HSet(hsetkey, lastTimekey, utils.TimeStamp())
	//sdk类型
	pipeline.HSet(hsetkey, sdkType, request.SdkType)
	//在线设备唯一ID
	pipeline.HSet(hsetkey, deviceKey, request.GetDeviceId())

	//状态key
	serverIpKey := constants.GetUserAccessIpKey(strUserId, request.GetTType())
	clientIpKey := constants.GetUserClientIpKey(strUserId, request.GetTType())

	//服务器接入层Ip
	pipeline.Set(serverIpKey, request.GetServerIp(), time.Second*time.Duration(request.GetNextRequest())+constants.DELAY_SECOND_TIME)
	//客户端IP
	pipeline.Set(clientIpKey, request.GetClientIp(), time.Second*time.Duration(request.GetNextRequest())+constants.DELAY_SECOND_TIME)

	//执行redis查询
	_, e := pipeline.Exec()
	if e != nil {
		return e
	}
	return e
}

//查询用户在线状态
func GetUserStatus(userId int32) (*dto.User_terminal_status, error) {
	pipeline := Pipeline()
	defer pipeline.Close()
	stringCmds := make([]*redis.StringCmd, 8)

	strUserId := fmt.Sprint(userId)

	//用户状态hset key
	hsetkey := constants.GetUserStatusHSetKey(strUserId)
	//用户SDK 类型
	sdkType := constants.GetSdkTypeKey()
	//移动端设备ID
	deviceKey := constants.GetUserDeviceKey(protobuffer.TerminalType_MOBILE)
	//最后一次在线时间
	lastTimekey := constants.GetUserLastTimeKey()
	stringCmds[0] = pipeline.HGet(hsetkey, lastTimekey)
	stringCmds[1] = pipeline.HGet(hsetkey, sdkType)
	stringCmds[2] = pipeline.HGet(hsetkey, deviceKey)

	//移动端
	mobileStatusKey := constants.GetUserAccessIpKey(strUserId, protobuffer.TerminalType_MOBILE)
	//pc端
	pcStatusKey := constants.GetUserAccessIpKey(strUserId, protobuffer.TerminalType_PC)
	//web端
	webStatusKey := constants.GetUserAccessIpKey(strUserId, protobuffer.TerminalType_WEB)
	stringCmds[3] = pipeline.Get(mobileStatusKey)
	stringCmds[4] = pipeline.Get(pcStatusKey)
	stringCmds[5] = pipeline.Get(webStatusKey)

	//长连接服务器地址
	mobileServerIpKey := constants.GetUserAccessIpKey(strUserId, protobuffer.TerminalType_MOBILE)
	mobileClientIpKey := constants.GetUserClientIpKey(strUserId, protobuffer.TerminalType_MOBILE)
	stringCmds[6] = pipeline.Get(mobileServerIpKey)
	stringCmds[7] = pipeline.Get(mobileClientIpKey)

	//执行redis查询
	_, e := pipeline.Exec()
	if e != nil && e.Error() != "redis: nil" {
		return nil, e
	}
	result := &dto.User_terminal_status{}
	//最后一次心跳时间
	if stringCmds[0] != nil {
		result.LastTime = stringCmds[0].Val()
	}
	//上次登录设备类型 0:手机，1:PC, 2:WEB 3:功能机不带拓展功能, 4:功能机带拓展功能(离线情况下为该字段代表上一次登录sdk类型)
	if stringCmds[1] != nil {
		r := stringCmds[1].Val()
		if len(r) > 0 {
			sdkType, _ := strconv.ParseInt(r, 10, 32)
			result.SdkType = int32(sdkType)
		}
	}
	//移动端设备ID
	if stringCmds[2] != nil {
		result.Mobile_DeviceId = stringCmds[2].Val()
	}
	//移动端是否在线
	if stringCmds[3] != nil {
		r := stringCmds[3].Val()
		if len(r) > 0 {
			result.Mobile_Status = true
		}
	}
	//pc端是否在线
	if stringCmds[4] != nil {
		r := stringCmds[4].Val()
		if len(r) > 0 {
			result.PC_Status = true
		}
	}
	//web端是否在线
	if stringCmds[5] != nil {
		r := stringCmds[5].Val()
		if len(r) > 0 {
			result.Web_Status = true
		}
	}
	//移动端长连接服务器IP
	if stringCmds[6] != nil {
		result.Mobile_AccessIp = stringCmds[6].Val()
	}
	//移动端IP
	if stringCmds[7] != nil {
		result.Mobile_ClientIp = stringCmds[7].Val()
	}
	return result, nil
}
*/
