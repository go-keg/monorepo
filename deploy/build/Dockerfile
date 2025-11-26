FROM golang:1.25 AS builder

ARG VERSION
ARG SERVICE

COPY . /src
WORKDIR /src
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go build -ldflags "-X main.Version=${VERSION}" -o ./bin/${SERVICE} ./cmd/${SERVICE}/

ARG IMAGE=ubuntu:latest
FROM ${IMAGE}

ARG SERVICE

COPY --from=builder /src/bin/${SERVICE} /app/${SERVICE}
COPY --from=builder /src/configs/${SERVICE}.yaml /app/configs/${SERVICE}.yaml
RUN mkdir "/app/logs"

WORKDIR /app

ENTRYPOINT ["/app/${SERVICE}"]