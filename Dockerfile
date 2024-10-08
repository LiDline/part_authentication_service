FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

RUN ls -l /app

EXPOSE 5000

CMD ["./test"]
