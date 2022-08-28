#build stage
FROM golang:alpine AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
