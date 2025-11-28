
FROM golang:1.24-alpine AS builder

LABEL authors="yannvanneth"

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./

RUN go build -o product-service ./cmd/server

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/product-service .

EXPOSE 9006

# Run the app
CMD ["./product-service"]
