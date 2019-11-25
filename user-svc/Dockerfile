FROM golang:1.13-alpine as build

WORKDIR /src

ADD . /src/

ARG GOPROXY=https://goproxy.io
ARG GOOS=linux
ARG GOARCH=amd64

RUN GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH GO111MODULE=on \
    go build -o /app/user-grpc github.com/ginuerzh/k8s-envoy-grpc-lb/user-svc/grpc && \
    GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH GO111MODULE=on \
    go build -o /app/user-gw github.com/ginuerzh/k8s-envoy-grpc-lb/user-svc/gw && \
    echo 'nobody:x:65534:' > /src/group.nobody && \
    echo 'nobody:x:65534:65534::/:' > /src/passwd.nobody && \
    GRPC_HEALTH_PROBE_VERSION=v0.3.1 && \
   wget -q -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM scratch

ENV GRPC_PORT=8000
ENV GW_PORT=8080

EXPOSE 8000
EXPOSE 8080

COPY --from=build /src/group.nobody /etc/group
COPY --from=build /src/passwd.nobody /etc/passwd
USER nobody:nobody

COPY --from=build /app/user-grpc /bin/user-grpc
COPY --from=build /app/user-gw /bin/user-gw
# Add grpc-health-probe to use with readiness and liveness probes
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe

ENTRYPOINT ["/bin/user-grpc"]
