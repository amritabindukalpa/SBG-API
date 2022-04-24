FROM golang:1.16-alpine

RUN mkdir -p /app
WORKDIR /app
ADD . /app

RUN go build ./main.go

CMD ["./main"]