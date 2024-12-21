FROM golang:1.23.2

WORKDIR /app

COPY ./ ./

# build app
RUN go mod download && go mod verify
RUN go build -o game-store ./cmd/app/main.go

CMD ["./game-store"]