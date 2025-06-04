Сборка докера

docker-compose up --build

Генерация для swagger

swag init -g ./cmd/app/main.go -o api/docs
