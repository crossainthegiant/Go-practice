package main

import (
	"context"
	"fmt"
	pb "github.com/crossainthegiant/Go-practice/gRPC/route"
	"google.golang.org/grpc"
	"log"
	"time"
)

func runFirst(client pb.RouteGuideClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	r, err := client.GetFeature(ctx, &pb.Point{Latitude: 310235000, Longitude: 121437403})
	if err != nil {
		log.Fatalf("call service failed: %v", err)
	}
	fmt.Println("call service success: ", r)
}

func main() {
	// 发起连接，WithInsecure表示使用不安全的连接，即不使用SSL,WithBlock就是把代码变成blocking，拨号成功才会往下走
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	defer conn.Close()

	c := pb.NewRouteGuideClient(conn)

	runFirst(c)

}
