package server

/**
* @Time       : 2020/11/22 9:07 下午
* @Author     : xumamba
* @Description: stream example server
 */
import (
	"context"

	"github.com/golang/protobuf/proto"

	pb "go-life/example/grpc/stream"
)

type server struct {
	pb.UnimplementedRouteGuideServer

	FeaturesDB []*pb.Feature
}

func (s *server) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.FeaturesDB{
		if proto.Equal(feature.Location, point){
			return feature, nil
		}
	}
	return &pb.Feature{Location: point}, nil
}

func (s *server) ListFeatures(rectangle *pb.Rectangle, featuresServer pb.RouteGuide_ListFeaturesServer) error {
	panic("implement me")
}

func (s *server) RecordRoute(routeServer pb.RouteGuide_RecordRouteServer) error {
	panic("implement me")
}

