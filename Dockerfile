FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./pkg/service/service_test.go

RUN go build -o main ./cmd/main.go

CMD ["./main"]