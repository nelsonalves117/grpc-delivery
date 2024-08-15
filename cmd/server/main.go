package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"time"

	"github.com/nelsonalves117/gRPC-delivery/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Delivery struct {
	pb.UnimplementedDeliveryServer
}

func main() {
	addr := ":8080"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}

	srv := createServer()
	log.Printf("info: server ready on port %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

func createServer() *grpc.Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(timingInterceptor))

	var u Delivery
	pb.RegisterDeliveryServer(srv, &u)

	reflection.Register(srv)

	return srv
}

func timingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)
		log.Printf("info: %s took %v", info.FullMethod, duration)
	}()

	return handler(ctx, req)
}

func (d *Delivery) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "no metadata found")
	}
	log.Printf("info: api_key %s", md["api_key"])

	resp := pb.StartResponse{
		OrderId: req.OrderId,
	}

	return &resp, nil
}

func (d *Delivery) End(ctx context.Context, req *pb.EndRequest) (*pb.EndResponse, error) {
	resp := pb.EndResponse{
		OrderId: req.OrderId,
	}

	return &resp, nil
}

func (d *Delivery) Location(stream pb.Delivery_LocationServer) error {
	count := int64(0)
	driverId := ""

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status.Error(codes.Internal, "can't read")
		}

		driverId = req.DriverId
		count++
	}

	resp := pb.LocationResponse{
		DriverId: driverId,
		Count:    count,
	}

	return stream.SendAndClose(&resp)
}
