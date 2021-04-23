//



package utils

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"net"
	"strings"
)
/*
func ConvertNodebIdListToProtoMessageList(l []*entities.NbIdentity) []proto.Message {
	protoMessageList := make([]proto.Message, len(l))

	for i, d := range l {
		protoMessageList[i] = d
	}

	return protoMessageList
}
*/
func MarshalProtoMessageListToJsonArray(msgList []proto.Message) (string, error){
	m := jsonpb.Marshaler{}
	ms := "["

	for _, msg := range msgList {
		s, err :=m.MarshalToString(msg)

		if (err != nil) {
			return s, err;
		}


		ms+=s+","
	}

	return strings.TrimSuffix(ms,",") +"]", nil
}

func GetValue(key string, defaultvalue string) string {
	if key == "" {
		return defaultvalue
	}
	return key
}

func GetHostIp() (string ,error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(),nil
			}
		}
	}
	return "", fmt.Errorf("Can not find the host ip address!")

}