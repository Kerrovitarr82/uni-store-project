FROM golang:1.23.2

WORKDIR /app

COPY ./ ./

# build app
RUN go mod download && go mod verify

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN $(go env GOPATH)/bin/swag init -g ./cmd/app/main.go -o api/docs

RUN go build -o game-store ./cmd/app/main.go

CMD ["./game-store"]