package constants

import "strconv"

const baseKey  = "ss"  //status-server
const LastTimeKeySubKey  = "l"
const NextHeartbeatIntervalSubKey  = "n"
const RegisterInfoSubKey  = "i"

//取得hmap key ss:{identity}
func GetIdentityKey(identity string) string {
	return baseKey + Colon + identity
}

//最后一次心跳时间 ｛deviceType｝:l
func GetLastTimeKey(deviceType byte) string {
	return strconv.Itoa(int(deviceType)) + Colon + LastTimeKeySubKey
}
//心跳间隔key  ｛deviceType｝:n
func GetNextHeartbeatIntervalKey(deviceType byte) string {
	return strconv.Itoa(int(deviceType)) + Colon + NextHeartbeatIntervalSubKey
}

//注册信息key  ｛deviceType｝:i
func GetRegisterInfoKey(deviceType byte) string {
	return strconv.Itoa(int(deviceType)) + Colon + RegisterInfoSubKey
}