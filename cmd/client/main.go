package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nelsonalves117/gRPC-delivery/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	addr := "localhost:8080"
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer conn.Close()

	log.Printf("info: connected to %s", addr)
	c := pb.NewDeliveryClient(conn)

	req := pb.StartRequest{
		OrderId:      "order123",
		CustomerId:   "customer456",
		RestaurantId: "restaurant789",
		DriverId:     "driver012",
		DeliveryLocation: &pb.Location{
			Latitude:  37.7749,
			Longitude: -122.4194,
		},
		ItemIds:   []string{"item1", "item2", "item3"},
		OrderTime: timestamppb.Now(),
		Status:    pb.OrderStatus_ACCEPTED,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "api_key", "s3cr3t_k3y")

	resp, err := c.Start(ctx, &req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(resp)

	ereq := pb.EndRequest{
		OrderId:      "order123",
		DeliveryTime: timestamppb.Now(),
		TotalAmount:  100,
	}

	eresp, err := c.End(ctx, &ereq)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(eresp)

	stream, err := c.Location(ctx)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	lreq := pb.LocationRequest{
		DriverId: "007",
		Location: &pb.Location{
			Latitude:  51.4871871,
			Longitude: -0.1266743,
		},
	}

	for i := 0.000; i < 0.010; i += 0.001 {
		fmt.Printf("Latitude: %f\nLongitude: %f\n\n", lreq.Location.Latitude, lreq.Location.Longitude)
		lreq.Location.Latitude += 1
		lreq.Location.Longitude += 1
		if err := stream.Send(&lreq); err != nil {
			log.Fatalf("error: %s", err)
		}
	}
	lresp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(lresp)

}
