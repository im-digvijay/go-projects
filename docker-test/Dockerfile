FROM golang:1.21.4-alpine3.18

RUN mkdir /app

COPY go.mod /app/go.mod
COPY main.go /app/main.go

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]