package main

import (
	"fmt"
	"log"

	"github.com/nelsonalves117/gRPC-delivery/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
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
	fmt.Println(&req)

	data, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var req2 pb.StartRequest
	if err := proto.Unmarshal(data, &req2); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(&req2)
}
