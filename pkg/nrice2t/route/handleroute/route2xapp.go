package handleroute

import (
	"fmt"
	pbe2t "nRIC/api/v1/pb/nrice2t"
	"nRIC/internal/msgx"
)

type Route2Xapp struct {
	Route map[int64]*pbe2t.RouteTable
	Endpoint map[int64]*msgx.KafkaMsgSender
}

func NewRoute2Xapp() *Route2Xapp {
	return &Route2Xapp{
		Route:    make(map[int64]*pbe2t.RouteTable),
		Endpoint: make(map[int64]*msgx.KafkaMsgSender),
	}
}


func Drop() {
}


func (rt *Route2Xapp)Insert(route *pbe2t.RouteTable) (*pbe2t.RouteTableInsertResponse, error){
	fmt.Printf("Insert SubIdXapp = %d,Topic = %s,SubIdRan = %d\n",route.SubIdXapp,route.Topic,route.SubIdRan)
	rt.Route[route.SubIdXapp] = route
	_,ok := rt.Endpoint[route.SubIdXapp]
	if !ok {
		rt.Endpoint[route.SubIdXapp] = msgx.NewKafkaMsgSender(route.Topic)
	}


	return &pbe2t.RouteTableInsertResponse{SubIdXapp: route.SubIdXapp},nil
}


func (rt *Route2Xapp) Update(route *pbe2t.RouteTable) (*pbe2t.RouteTableUpdateResponse, error){
	fmt.Printf("Update SubIdXapp = %d,Topic = %s,SubIdRan = %d\n",route.SubIdXapp,route.Topic,route.SubIdRan)
	rt.Route[route.SubIdXapp] = route
	_,ok := rt.Endpoint[route.SubIdXapp]
	if !ok {
		rt.Endpoint[route.SubIdXapp] = msgx.NewKafkaMsgSender(route.Topic)
	}
	return &pbe2t.RouteTableUpdateResponse{Updated: route.SubIdXapp},nil
}


func (rt *Route2Xapp) List() ([]*pbe2t.RouteTable,error) {

	return nil,nil
}

func (rt *Route2Xapp) GetbySubIdXapp(SubIdXapp int64) (*pbe2t.RouteTable,error) {
	return rt.Route[SubIdXapp], nil
}

func (rt *Route2Xapp) GetbySubIdRan(SubIdRan int64) ([]*pbe2t.RouteTable,error) {
	routelist := []*pbe2t.RouteTable{}
	for _,r := range rt.Route {
		if r.SubIdRan == SubIdRan {
			routelist = append(routelist,r)
		}
	}
	return routelist, nil
}

func (rt *Route2Xapp) GetEndpointbySubIdRan(SubIdRan int64) (map[int64]*msgx.KafkaMsgSender,error) {
	Endpointmap := make(map[int64]*msgx.KafkaMsgSender)
	//routelist := []*pbe2t.RouteTable{}
	for _,r := range rt.Route {
		if r.SubIdRan == SubIdRan {
			Endpointmap[r.SubIdXapp] = rt.Endpoint[r.SubIdXapp]
		}
	}
	return Endpointmap, nil
}