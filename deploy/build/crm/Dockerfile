
ARG IMAGE=ubuntu:latest
FROM ${IMAGE}

WORKDIR /app
COPY bin/crm /app/crm
COPY configs/crm.yaml /app/configs/crm.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/crm"]