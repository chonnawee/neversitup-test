FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY . .

RUN go get -d -v ./...

EXPOSE 8080

CMD ["go", "run", "./cmd/server/main.go"]