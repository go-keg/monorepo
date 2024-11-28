FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates curl vim \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y 