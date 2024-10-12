FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

FROM scratch

COPY --from=build /app/server /

ENTRYPOINT [ "/server" ]