FROM golang:1.21.4-bookworm AS builder

WORKDIR /app
COPY . /app
RUN go build cmd/ip/firewall/addrlists/main.go

FROM debian:bookworm-slim

RUN set -eux; \
    apt update; \
    apt install -y ca-certificates; \
    rm -rf /var/cache/apt/archives /var/lib/apt/lists/*

ENV ADDRESS_LIST_URL='https://github.com/17mon/china_ip_list/raw/master/china_ip_list.txt'

WORKDIR /app
COPY --from=builder /app/main /app/

CMD /app/main
