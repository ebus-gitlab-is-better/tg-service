FROM golang:1.19 AS builder

COPY . /src
WORKDIR /src

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY  ./configs/config.yaml /app

WORKDIR /app

EXPOSE 9000

CMD ["./tgbot-service", "-conf", "config.yaml"]
