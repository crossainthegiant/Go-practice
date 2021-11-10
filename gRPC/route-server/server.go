package main

import (
	"context"
	"fmt"
	pb "github.com/crossainthegiant/Go-practice/gRPC/route"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	features []*pb.Feature
}

//实现接口
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		} //比较两个message是否相同
	}

	return nil, nil
}

func (s *routeGuideServer) ListFeatures(*pb.Rectangle, pb.RouteGuide_ListFeaturesServer) error {
	return nil
}

// client side streaming,与上面相反
func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

//bidirectional streaming，都是流
func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{Name: "上海交大", Location: &pb.Point{
				Latitude:  310235000,
				Longitude: 121437403,
			}},
		},
	}
}

func main() {
	lisener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the addr")
	}
	//grpc服务器
	grpcServer := grpc.NewServer()
	//将服务器与处理器绑定
	pb.RegisterRouteGuideServer(grpcServer, newServer())

	//reflection.Register(s)
	fmt.Println("gRPC server listen in 5000...")
	err = grpcServer.Serve(lisener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
