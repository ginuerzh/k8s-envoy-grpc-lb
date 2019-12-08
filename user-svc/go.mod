module github.com/ginuerzh/k8s-envoy-grpc-lb/user-svc

go 1.13

require (
	github.com/ginuerzh/echo v0.0.1
	github.com/ginuerzh/k8s-envoy-grpc-lb/account-svc v0.0.0-20191204114315-ea51ecbc1b19 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/opentracing/opentracing-go v1.1.0
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc v1.25.1
)
