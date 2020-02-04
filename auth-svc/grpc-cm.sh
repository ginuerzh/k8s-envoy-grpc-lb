#!/bin/sh

kubectl create configmap auth-grpc-config --from-literal=grpc-addr=:8000 --from-literal=redis-url=redis://localhost:6379/0 --from-literal=ttl=300s
