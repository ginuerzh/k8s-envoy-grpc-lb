//go:generate protoc -I../api --grpc-gateway_out=logtostderr=true:../api ../api/user.proto

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	api "github.com/ginuerzh/k8s-envoy-grpc-lb/user-svc/api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func run() error {
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "8000"
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := api.RegisterUserHandlerFromEndpoint(ctx, mux, ":"+grpcPort, opts); err != nil {
		return err
	}

	gwPort := os.Getenv("GW_PORT")
	if gwPort == "" {
		gwPort = "8080"
	}
	log.Printf("listening on port %s forward to %s", gwPort, grpcPort)

	return http.ListenAndServe(":"+gwPort, mux)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
