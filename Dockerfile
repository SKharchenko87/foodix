FROM golang:alpine3.22 AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/foodix ./cmd/server/main.go

FROM alpine:3.22
WORKDIR /app

COPY --from=builder /build/bin/foodix .

ENTRYPOINT ["/app/foodix"]