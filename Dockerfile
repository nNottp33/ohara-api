FROM golang:1.25-alpine AS builder

RUN apk add --no-cache make bash

WORKDIR /app

COPY . .

RUN go mod tidy && go mod download

RUN make build

FROM alpine:3.14 AS runner

RUN apk add --no-cache ca-certificates tzdata

ENV TZ="Asia/Bangkok"

WORKDIR /app

COPY --from=builder /bin/app .

ENTRYPOINT ["./app"]