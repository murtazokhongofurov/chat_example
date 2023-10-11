FROM golang:1.20.0-alpine3.16 AS builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o main cmd/main.go

RUN apk add upx

RUN upx main

FROM scratch

WORKDIR /app

COPY --from=builder /app .

EXPOSE 1111

CMD ["/app/main"]