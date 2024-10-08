FROM ubuntu
MAINTAINER eiixy <eiixyw@gmail.com>

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl vim

WORKDIR /app
COPY bin/admin /app/admin
COPY configs/admin.yaml /app/configs/admin.yaml
RUN mkdir "/app/logs"

ENTRYPOINT ["/app/admin"]
