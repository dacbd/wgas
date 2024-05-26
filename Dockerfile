FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN go build

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /app/gaa /usr/local/bin/gaa
ENTRYPOINT ["gaa"]
