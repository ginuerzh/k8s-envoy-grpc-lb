#!/bin/sh

kubectl create cm envoy-user-config --from-file=envoy.yaml -n grpc-lb -o yaml --dry-run | kubectl replace -f -
