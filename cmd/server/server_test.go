package main

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/nelsonalves117/gRPC-delivery/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestEnd(t *testing.T) {
	req := pb.EndRequest{
		OrderId:      "order123",
		DeliveryTime: timestamppb.Now(),
		TotalAmount:  100,
	}

	var srv Delivery
	resp, err := srv.End(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.OrderId != req.OrderId {
		t.Fatalf("bad response id: got %#v, expected %#v", resp.OrderId, req.OrderId)
	}

}

func TestEnd2End(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	srv := createServer()
	go srv.Serve(lis)

	port := lis.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("localhost:%d", port)

	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Fatal(err)
	}

	c := pb.NewDeliveryClient(conn)

	req := pb.EndRequest{
		OrderId:      "order123",
		DeliveryTime: timestamppb.Now(),
		TotalAmount:  100,
	}

	resp, err := c.End(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.OrderId != req.OrderId {
		t.Fatalf("bad response id: got %#v, expected %#v", resp.OrderId, req.OrderId)
	}

}
