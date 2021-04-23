package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/pkg/dbagent/grpcserver/endpoint"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////

type msgServer struct {
	grpcMsgService          		grpctransport.Handler
	routeTableCreate 				grpctransport.Handler
	routeTableRead					grpctransport.Handler
	routeTableUpdate				grpctransport.Handler
	routeTableDelete				grpctransport.Handler
	routeTableReadAll				grpctransport.Handler

	ranFunctionsTableCreate 		grpctransport.Handler
	ranFunctionsTableRead			grpctransport.Handler
	ranFunctionsTableUpdate			grpctransport.Handler
	ranFunctionsTableDelete			grpctransport.Handler
	ranFunctionsTableReadAll		grpctransport.Handler

	moiTableCreate 					grpctransport.Handler
	moiTableRead					grpctransport.Handler
	moiTableUpdate					grpctransport.Handler
	moiTableDelete					grpctransport.Handler
	moiTableReadAll					grpctransport.Handler

}

//TBD  pbe2t need change to pbdb
func NewmsgServer(ep endpoint.MsgServiceEndpoint) pbdb.DbServiceServer {
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
		ranFunctionsTableCreate: grpctransport.NewServer(
			ep.RANFunctionsTableInsertEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		ranFunctionsTableRead: grpctransport.NewServer(
			ep.RANFunctionsTableReadEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		ranFunctionsTableUpdate: grpctransport.NewServer(
			ep.RANFunctionsTableUpdateEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		ranFunctionsTableDelete: grpctransport.NewServer(
			ep.RANFunctionsTableDeleteEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		ranFunctionsTableReadAll: grpctransport.NewServer(
			ep.RANFunctionsTableReadAllEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		moiTableCreate: grpctransport.NewServer(
			ep.MOITableInsertEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		moiTableRead: grpctransport.NewServer(
			ep.MOITableReadEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		moiTableUpdate: grpctransport.NewServer(
			ep.MOITableUpdateEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		moiTableDelete: grpctransport.NewServer(
			ep.MOITableDeleteEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
		moiTableReadAll: grpctransport.NewServer(
			ep.MOITableReadAllEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
	}
}

func (g *msgServer) RpcHandleMsg(ctx context.Context, r *pbdb.GrpcMsg) (*pbdb.GrpcReply, error) {
	_, rep, err := g.grpcMsgService.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.GrpcReply)
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

