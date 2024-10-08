FROM ubuntu
MAINTAINER eiixy <eiixyw@gmail.com>

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl vim

WORKDIR /app
COPY bin/api /app/api
COPY configs/api.yaml /app/configs/api.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/api"]
