//go:generate protoc -I../api --go_out=plugins=grpc:../api ../api/account.proto
//go:generate protoc -I../api --swagger_out=logtostderr=true:../api ../api/account.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	api "github.com/ginuerzh/k8s-envoy-grpc-lb/account-svc/api"
	"github.com/ginuerzh/k8s-envoy-grpc-lb/account-svc/svc/account"
	"github.com/ginuerzh/k8s-envoy-grpc-lb/account-svc/svc/health"
	userapi "github.com/ginuerzh/k8s-envoy-grpc-lb/user-svc/api"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8000"
	}

	userEndpoint := os.Getenv("USER_ENDPOINT")
	if userEndpoint == "" {
		userEndpoint = "user.grpc-lb:8000"
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening for user on port", port)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			unaryServerRecoveryInterceptor(),
			// unaryServerOpenTracingInterceptor(tracer),
			// unaryServerAuthInterceptor(),
			unaryServerLoggingInterceptor(),
		),
	)

	grpcConn, err := grpc.DialContext(context.Background(), userEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial svc1: %v", err)
	}

	api.RegisterAccountServer(s, &account.Server{
		UserClient: userapi.NewUserClient(grpcConn),
	})
	grpc_health_v1.RegisterHealthServer(s, &health.Server{})
	reflection.Register(s)

	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
