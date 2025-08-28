ARG IMAGE=ubuntu:latest
FROM ${IMAGE}

WORKDIR /app
COPY bin/account /app/account
COPY configs/account.yaml /app/configs/account.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/account"]
