FROM golang:1.20 AS builder

ARG VERSION
ARG SERVICE

COPY . /src
WORKDIR /src
RUN export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go build -ldflags "-X main.Version=${VERSION}" -o ./bin/${SERVICE} ./cmd/${SERVICE}/

FROM debian:stable-slim

ARG SERVICE

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin/${SERVICE} /app/service
COPY --from=builder /src/configs/${SERVICE}.yaml /app/configs/${SERVICE}.yaml
RUN mkdir "/app/logs"

WORKDIR /app

ENTRYPOINT ["/app/service"]