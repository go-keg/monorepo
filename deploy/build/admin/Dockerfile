ARG IMAGE=ubuntu:latest
FROM ${IMAGE}

WORKDIR /app
COPY bin/admin /app/admin
COPY configs/admin.yaml /app/configs/admin.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/admin"]
