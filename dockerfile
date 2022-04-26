FROM golang:1.16-alpine

WORKDIR /app
COPY . .

EXPOSE 8080

RUN go build ./main.go

CMD ["./main"]