ARG IMAGE=ubuntu:latest
FROM ${IMAGE}

WORKDIR /app
COPY bin/example /app/example
COPY configs/example.yaml /app/configs/example.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/example"]
