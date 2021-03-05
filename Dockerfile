FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir temp

RUN go build src/main.go

CMD ["./main"]
