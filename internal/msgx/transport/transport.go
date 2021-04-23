package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"nRIC/internal/msgx"
	"nRIC/internal/msgx/endpoint"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////

type msgServer struct {
	grpcMsgService          grpctransport.Handler
}

func NewMsgServer(ep endpoint.MsgServiceEndpoint) msgx.MsgServiceServer {
	return &msgServer{
		grpcMsgService: grpctransport.NewServer(
			ep.GrpcMsgServiceEndpoint,
			decodeGrpcRequest,
			encodeGrpcResponse,
		),
	}
}

func (g *msgServer) RpcHandleMsg(ctx context.Context, r *msgx.GrpcMsg) (*msgx.GrpcReply, error) {
	_, rep, err := g.grpcMsgService.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*msgx.GrpcReply)
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

