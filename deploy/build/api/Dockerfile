ARG IMAGE=ubuntu:latest
FROM ${IMAGE}

WORKDIR /app
COPY bin/api /app/api
COPY configs/api.yaml /app/configs/api.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/api"]
