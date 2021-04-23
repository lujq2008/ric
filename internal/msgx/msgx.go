package msgx

import "nRIC/internal/xapp"

type SenderIf interface {
	 SendMsg( *xapp.MsgParams) (error)
	 String() (string)
}

type EndpointList struct {
	EpList []SenderIf
}

func(l *EndpointList)AddEndpoint(S SenderIf) bool{
	l.EpList = append(l.EpList,S)
	return true
}

func(l *EndpointList)Size() int{
	return len(l.EpList)
}


func(l *EndpointList)HasEndpoint(S SenderIf) bool{
	for _,Ep := range l.EpList {
		if Ep == S {
			return true
		}
	}
	return false
}

func(l *EndpointList)String() string{
	return "EndpointList"
}