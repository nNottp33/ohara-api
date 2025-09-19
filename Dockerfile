FROM golang:1.25-alpine AS builder

RUN apk add --no-cache bash

WORKDIR /app

COPY . .

RUN go mod tidy && go mod download

RUN go build -o ./tmp/main ./cmd/server.go

FROM alpine:3.14 AS runner

RUN apk add --no-cache ca-certificates tzdata

ENV TZ="Asia/Bangkok"

WORKDIR /app

COPY --from=builder /app/tmp/main .

ENTRYPOINT ["./main"]
