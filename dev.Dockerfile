FROM golang:1.25-alpine

RUN apk add --no-cache git \
    && go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .

RUN mkdir -p tmp

CMD ["air", "-c", ".air.toml"]
