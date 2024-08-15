package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nelsonalves117/gRPC-delivery/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := runtime.NewServeMux()

	err = pb.RegisterDeliveryHandler(ctx, mux, conn)
	if err != nil {
		log.Fatal(err)
	}

	addr := ":8081"
	log.Printf("gateway server starting on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}

}
