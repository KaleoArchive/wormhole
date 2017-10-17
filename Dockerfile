FROM alpine:latest
COPY build/wormhole-linux-amd64 /usr/local/bin/wormhole
CMD ["/usr/local/bin/wormhole"]
