FROM golang:1.23

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build -o server cmd/main.go

CMD ["./server"]