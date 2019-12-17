#!/bin/sh

kubectl create cm envoy-auth-config --from-file=envoy.yaml -n grpc-lb -o yaml --dry-run | kubectl replace -f -
