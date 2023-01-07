# Build Stage
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY ../../go/go.mod go.sum ./

RUN go mod download

COPY ../../go .

RUN go build -ldflags "-w -s" -o server cmd/server/main.go


# Production Stage
FROM alpine:latest AS production

WORKDIR /app
RUN apk update

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]


# Local Stage
FROM golang:1.19 as devlop

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
