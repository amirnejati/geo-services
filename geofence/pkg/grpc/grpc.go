package grpc

import (
	"context"
	"github.com/amirnejati/map-services/geofence/pkg/endpoint"
	"github.com/amirnejati/map-services/geofence/pkg/grpc/pb"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	point2District kitgrpc.Handler
}

func NewGRPCServer(eps endpoint.EndpointsSet, options []kitgrpc.ServerOption) pb.GeofenceServer {
	return &grpcServer{
		point2District: kitgrpc.NewServer(
			eps.Point2DistrictEndpoint,
			decodeGRPCPoint2DistrictRequest,
			encodeGRPCPoint2DistrictResponse,
			options...,
		),
	}
}

func (g *grpcServer) Point2District(ctx context.Context, r *pb.Point2DistrictRequest) (*pb.Point2DistrictReply, error) {
	_, rep, err := g.point2District.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Point2DistrictReply), nil
}

func decodeGRPCPoint2DistrictRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Point2DistrictRequest)
	return endpoint.Point2DistrictRequest{Latitude: req.Latitude, Longitude: req.Longitude}, nil
}

func encodeGRPCPoint2DistrictResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoint.Point2DistrictResponse)
	dist := pb.Point2DistrictReply{
		PolygonId:  reply.District.PolygonId.Hex(),
		Name:       reply.District.Name,
		DistrictNo: uint32(reply.District.DistrictNo),
	}
	return &dist, nil
}
