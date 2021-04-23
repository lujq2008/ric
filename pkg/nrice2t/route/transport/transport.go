package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"nRIC/pkg/nrice2t/route/endpoint"
	pbe2t "nRIC/api/v1/pb/nrice2t"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////

type msgServer struct {
	grpcMsgService          grpctransport.Handler
	routeTableCreate 		grpctransport.Handler
	routeTableRead			grpctransport.Handler
	routeTableUpdate		grpctransport.Handler
	routeTableDelete		grpctransport.Handler
	routeTableReadAll		grpctransport.Handler
}

func NewmsgServer(ep endpoint.MsgServiceEndpoint) pbe2t.E2TServiceServer {
	return &msgServer{
		grpcMsgService: grpctransport.NewServer(
			ep.GrpcMsgServiceEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		routeTableCreate: grpctransport.NewServer(
			ep.RouteTableInsertEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		routeTableRead: grpctransport.NewServer(
			ep.RouteTableReadEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		routeTableUpdate: grpctransport.NewServer(
			ep.RouteTableUpdateEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		routeTableDelete: grpctransport.NewServer(
			ep.RouteTableDeleteEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		routeTableReadAll: grpctransport.NewServer(
			ep.RouteTableReadAllEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
	}
}

func (g *msgServer) RpcHandleMsg(ctx context.Context, r *pbe2t.GrpcMsg) (*pbe2t.GrpcReply, error) {
	_, rep, err := g.grpcMsgService.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbe2t.GrpcReply)
	return Resp, nil
}


func decodeGrpcRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	//req := grpcReq.(*GrpcMsg)
	return grpcReq, nil
}

func encodeGrpcResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	//reply := grpcReply.(*GrpcReply)
	return grpcReply, nil
}

