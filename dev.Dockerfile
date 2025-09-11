FROM golang:1.25-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .

RUN mkdir -p tmp

CMD ["air", "-c", ".air.toml"]
